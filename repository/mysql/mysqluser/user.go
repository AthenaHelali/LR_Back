package mysqluser

import (
	"database/sql"
	"fmt"
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/errormessage"
	"game-app/pkg/richerror"
	"game-app/repository/mysql"
	"log"
)

func (d *DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.conn.Connection().QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, richerror.New("mysql.IsPhoneNumberUnique").WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return false, nil
}
func (d *DB) RegisterUser(user entity.User) (entity.User, error) {
	res, err := d.conn.Connection().Exec(`insert into users(name, phone_number, password, role) values (?, ? , ?, ?)`, user.Name, user.PhoneNumber, user.Password, user.Role.String())
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command: %W", err)
	}

	id, _ := res.LastInsertId()
	user.ID = uint(id)

	return user, nil

}

func (d *DB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"
	row := d.conn.Connection().QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)

		}

		//TODO log unexpected error for better observability

		return entity.User{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return user, nil
}

func (d *DB) GetUserByID(UserID uint) (entity.User, error) {
	const op = "mysql.GetUserByID"
	row := d.conn.Connection().QueryRow(`select * from users where id = ?`, UserID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)

		}
		return entity.User{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return user, nil
}

func scanUser(scanner mysql.Scanner) (entity.User, error) {
	var user entity.User

	var roleStr string

	err := scanner.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.CreatedAt, &user.Password, &roleStr)

	user.Role = entity.MapToRoleEntity(roleStr)
	return user, err

}

func (d *DB) AddFavoriteLaptop(userID int, laptopID int) error {
	const op = "mysql.AddFavoriteLaptop"

	_, err := d.conn.Connection().Exec(`insert into user_laptop(user_ref, laptop_ref) values (?, ?)`, userID, laptopID)
	if err != nil {
		return fmt.Errorf("can't execute command: %W", err)
	}

	return nil
}

func (d *DB) RemoveFavoriteLaptop(LaptopID int, UserID int) error {
	const op = "mysql.DeleteLaptop"
	deleteQuery := "DELETE FROM user_laptop WHERE laptop_ref = ? and user_ref = ? "
	// Execute the delete query
	_, err := d.conn.Connection().Exec(deleteQuery, LaptopID, UserID)
	if err != nil {
		return richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
	}

	return nil

}

func (d *DB) GetLaptops(UserID uint) ([]entity.Laptop, error) {
	const op = "mysql.GetLaptops"
	var laptops []entity.Laptop
	rows, err := d.conn.Connection().Query(`select laptops.id, laptops.cpu, laptops.ram, laptops.ssd, laptops.hdd, laptops.graphic,laptops.screen_size, laptops.company,laptops.price, laptops.image_url, laptops.redirect_url from laptops join user_laptop on laptops.id = user_laptop.laptop_ref where user_laptop.user_ref = ?`, UserID)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	for rows.Next() {
		var laptop entity.Laptop
		var imageUrlTmp *string = new(string)
		var redirectUrlTmp *string = new(string)

		err := rows.Scan(&laptop.ID, &laptop.CPU, &laptop.RAM, &laptop.SSD, &laptop.HDD, &laptop.Graphic, &laptop.ScreenSize, &laptop.Company, &laptop.Price, &imageUrlTmp, &redirectUrlTmp)
		if err != nil {
			println(err.Error())
			return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}
		if imageUrlTmp == nil {
			laptop.ImageURL = ""
		} else {
			laptop.ImageURL = *imageUrlTmp
		}
		if redirectUrlTmp == nil {
			laptop.RedirectURL = ""
		} else {
			laptop.RedirectURL = *redirectUrlTmp
		}

		laptops = append(laptops, laptop)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return []entity.Laptop{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)

		}
		return []entity.Laptop{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return laptops, nil
}

func (d *DB) GetLaptopByID(LaoptopID uint) (entity.Laptop, error) {
	const op = "mysql.GetLaptops"
	row := d.conn.Connection().QueryRow(`select laptops.id, laptops.cpu, laptops.ram, laptops.ssd, laptops.hdd, laptops.graphic,laptops.screen_size, laptops.company,laptops.price, laptops.image_url, laptops.redirect_url from laptops where id = ?`, LaoptopID)
	if err := row.Err(); err != nil {
		return entity.Laptop{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	var laptop entity.Laptop
	var imageUrlTmp *string = new(string)
	var redirectUrlTmp *string = new(string)

	err := row.Scan(&laptop.ID, &laptop.CPU, &laptop.RAM, &laptop.SSD, &laptop.HDD, &laptop.Graphic, &laptop.ScreenSize, &laptop.Company, &laptop.Price, &imageUrlTmp, &redirectUrlTmp)
	if err != nil {
		println(err.Error())
		return laptop, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
	}
	if imageUrlTmp == nil {
		laptop.ImageURL = ""
	} else {
		laptop.ImageURL = *imageUrlTmp
	}
	if redirectUrlTmp == nil {
		laptop.RedirectURL = ""
	} else {
		laptop.RedirectURL = *redirectUrlTmp
	}

	return laptop, nil
}

func (d *DB) UpdateUser(updatedUser entity.User) error {

	// Create an update query based on non-empty field
	updateQuery := "UPDATE users SET"
	updateValues := []interface{}{}

	if updatedUser.Name != "" {
		updateQuery += " name = ?,"
		updateValues = append(updateValues, updatedUser.Name)
	}

	if updatedUser.Password != "" {
		updateQuery += " password = ?,"
		updateValues = append(updateValues, updatedUser.Password)
	}

	if updatedUser.PhoneNumber != "" {
		updateQuery += " phone_number = ?,"
		updateValues = append(updateValues, updatedUser.PhoneNumber)
	}

	// Remove the trailing comma
	updateQuery = updateQuery[:len(updateQuery)-1]

	updateQuery += " WHERE id = ?"
	updateValues = append(updateValues, updatedUser.ID)

	fmt.Print(updateQuery)

	// Execute the update query
	_, err := d.conn.Connection().Exec(updateQuery, updateValues...)
	return err
}

func (d *DB) Search(IDs []int) ([]param.LaptopInfo, error) {
	const op = "mysql.Search"
	var laptops []param.LaptopInfo
	var info param.LaptopInfo

	for _, id := range IDs {
		row := d.conn.Connection().QueryRow(`select laptops.id, laptops.cpu, laptops.ram, laptops.ssd, laptops.hdd, laptops.graphic,laptops.screen_size, laptops.company,laptops.price, laptops.image_url, laptops.redirect_url from laptops where id = ?`, id+1)
		if err := row.Err(); err != nil {
			return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
		}
		err := row.Scan(&info.ID, &info.CPU, &info.RAM, &info.SSD, &info.HDD, &info.Graphic, &info.ScreenSize, &info.Company, &info.Price, &info.ImageURL, &info.RedirectURL)
		if err != nil {
			return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
		}
		laptops = append(laptops, info)

	}

	return laptops, nil
}

func (d *DB) AddLaptop(LaptopInfo param.LaptopInfo, UserID uint) (uint, error) {
	const op = "mysql.AddLaptop"

	res, err := d.conn.Connection().Exec(`insert into laptops(cpu, ram, ssd, hdd, graphic, screen_size, company, price, created_at, image_url) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, LaptopInfo.CPU, LaptopInfo.RAM, LaptopInfo.SSD, LaptopInfo.HDD, LaptopInfo.Graphic, LaptopInfo.ScreenSize, LaptopInfo.Company, LaptopInfo.Price, LaptopInfo.CreatedAt, LaptopInfo.ImageURL)
	if err != nil {
		return 0, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return uint(lastInsertID), nil

}

func (d *DB) AddSellerLaptop(LaptopID, UserID uint) error {
	_, err := d.conn.Connection().Exec(`insert into seller_laptop(seller_ref, laptop_ref) values (?, ?)`, UserID, LaptopID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("can't execute command: %W", err)
	}

	return nil
}

func (d *DB) GetSellerLaptops(UserID uint) ([]entity.Laptop, error) {
	const op = "mysql.GetLaptops"
	var laptops []entity.Laptop
	rows, err := d.conn.Connection().Query(`select laptops.id, laptops.cpu, laptops.ram, laptops.ssd, laptops.hdd, laptops.graphic,laptops.screen_size, laptops.company,laptops.price, laptops.image_url, laptops.redirect_url from laptops join seller_laptop on laptops.id = seller_laptop.laptop_ref where seller_laptop.seller_ref = ?`, UserID)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	for rows.Next() {
		var laptop entity.Laptop
		var imageUrlTmp *string = new(string)
		var redirectUrlTmp *string = new(string)

		err := rows.Scan(&laptop.ID, &laptop.CPU, &laptop.RAM, &laptop.SSD, &laptop.HDD, &laptop.Graphic, &laptop.ScreenSize, &laptop.Company, &laptop.Price, &imageUrlTmp, &redirectUrlTmp)
		if err != nil {
			println(err.Error())
			return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}
		if imageUrlTmp == nil {
			laptop.ImageURL = ""
		} else {
			laptop.ImageURL = *imageUrlTmp
		}
		if redirectUrlTmp == nil {
			laptop.RedirectURL = ""
		} else {
			laptop.RedirectURL = *redirectUrlTmp
		}

		laptops = append(laptops, laptop)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return []entity.Laptop{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)

		}
		return []entity.Laptop{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return laptops, nil
}

func (d *DB) UpdateLaptop(updatedLaptop entity.Laptop) error {

	// Create an update query based on non-empty field
	updateQuery := "UPDATE laptops SET"
	updateValues := []interface{}{}

	if updatedLaptop.CPU != "" {
		updateQuery += " cpu = ?,"
		updateValues = append(updateValues, updatedLaptop.CPU)
	}

	if updatedLaptop.RAM != 0 {
		updateQuery += " ram = ?,"
		updateValues = append(updateValues, updatedLaptop.RAM)
	}

	if updatedLaptop.SSD != 0 {
		updateQuery += " ssd = ?,"
		updateValues = append(updateValues, updatedLaptop.SSD)
	}

	if updatedLaptop.HDD != 0 {
		updateQuery += " hdd = ?,"
		updateValues = append(updateValues, updatedLaptop.HDD)
	}

	if updatedLaptop.Graphic != 0 {
		updateQuery += " graphic = ?,"
		updateValues = append(updateValues, updatedLaptop.Graphic)
	}

	if updatedLaptop.ScreenSize != "" {
		updateQuery += " screen_size = ?,"
		updateValues = append(updateValues, updatedLaptop.ScreenSize)
	}

	if updatedLaptop.Company != "" {
		updateQuery += " company = ?,"
		updateValues = append(updateValues, updatedLaptop.Company)
	}

	if updatedLaptop.Price != "" {
		updateQuery += " price = ?,"
		updateValues = append(updateValues, updatedLaptop.Price)
	}
	if updatedLaptop.RedirectURL != "" {
		updateQuery += " redirect_url = ?,"
		updateValues = append(updateValues, updatedLaptop.RedirectURL)
	}

	// Remove the trailing comma
	updateQuery = updateQuery[:len(updateQuery)-1]

	updateQuery += " WHERE id = ?"
	updateValues = append(updateValues, updatedLaptop.ID)

	_, err := d.conn.Connection().Exec(updateQuery, updateValues...)
	return err
}

func (d *DB)RemoveSellerLaptop(LaptopID int, SellerID int) error{
	const op = "mysql.DeleteLaptop"
	deleteQuery := "DELETE FROM seller_laptop WHERE laptop_ref = ? and seller_ref = ? "
	_, err := d.conn.Connection().Exec(deleteQuery, LaptopID, SellerID)
	if err != nil {
		return richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
	}

	deleteQuery = "DELETE FROM laptops WHERE  id = ?"
	_, err = d.conn.Connection().Exec(deleteQuery, LaptopID)
	if err != nil {
		return richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
	}

	return nil

}
