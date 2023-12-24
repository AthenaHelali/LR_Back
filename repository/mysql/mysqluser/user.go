package mysqluser

import (
	"database/sql"
	"fmt"
	"game-app/entity"
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

func (d *DB) AddFavoriteLaptop(laptop entity.Laptop, userID int) (entity.Laptop, error) {
	const op = "mysql.AddFavoriteLaptop"


//after merging databases, laptop already exists in database and just execute second query
	res, err := d.conn.Connection().Exec(`insert into laptops(cpu, ram,ssd, hdd, graphic, screen_size, company, price, image_url, redirect_url ) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, laptop.CPU, laptop.RAM, laptop.SSD, laptop.HDD, laptop.Graphic, laptop.ScreenSize, laptop.Company,laptop.Price, laptop.ImageURL, laptop.RedirectURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Laptop{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)

		}
		return entity.Laptop{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	id, _ := res.LastInsertId()
	laptop.ID = uint(id)

	_, err = d.conn.Connection().Exec(`insert into user_laptop(user_ref, laptop_ref) values (?, ?)`, userID, laptop.ID)
	if err != nil {
		return entity.Laptop{}, fmt.Errorf("can't execute command: %W", err)
	}

	return laptop, nil
}

func (d *DB)RemoveFavoriteLaptop(LaptopID int, UserID int)error{
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
		err := rows.Scan(&laptop.ID, &laptop.CPU, &laptop.RAM, &laptop.SSD, &laptop.HDD, &laptop.Graphic, &laptop.ScreenSize, &laptop.Company,&laptop.Price, &laptop.ImageURL, &laptop.RedirectURL)
		if err != nil {
			return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
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
	row := d.conn.Connection().QueryRow(`select laptops.id, laptops.cpu, laptops.ram, laptops.ssd, laptops.hdd, laptops.graphic,laptops.screen_size, laptops.company,laptops.price, laptops.image_url, laptops.redirect_url from laptops where laptop.ID = ?`, LaoptopID)
	if err := row.Err(); err!= nil {
		return entity.Laptop{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	
	var laptop entity.Laptop
	err := row.Scan(&laptop.ID, &laptop.CPU, &laptop.RAM, &laptop.SSD, &laptop.HDD, &laptop.Graphic, &laptop.ScreenSize, &laptop.Company,&laptop.Price, &laptop.ImageURL, &laptop.RedirectURL)
	if err != nil {
		return entity.Laptop{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return laptop, nil
}

func (db *DB)UpdateUser(updatedUser entity.User) error {

	// Create an update query based on non-empty fields
	fmt.Println(updatedUser.Name)
	fmt.Println(updatedUser.Password)
	fmt.Println(updatedUser.Name)
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
	_, err := db.conn.Connection().Exec(updateQuery, updateValues...)
	return err
}

