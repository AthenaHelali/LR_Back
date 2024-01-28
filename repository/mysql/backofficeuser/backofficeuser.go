package mysqlbackofficeuser

import (
	"game-app/entity"
	"game-app/pkg/errormessage"
	"game-app/pkg/richerror"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (d *DB) ListAllUsers() ([]entity.User, error) {
	const op = "mysql.ListAllUsers"
	rows, err := d.conn.Connection().Query("SELECT id,name, phone_number, created_at, password, role FROM users")
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)

	}
	defer rows.Close()

	// Create a slice to hold the users
	var users []entity.User

	// Iterate over the result set and populate the users slice
	for rows.Next() {
		var user entity.User
		var role string
		err := rows.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.CreatedAt, &user.Password, &role)
		if err != nil {
			return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)

		}

		if role == "user" {
			user.Role = 1
		} else if role == "admin" {
			user.Role = 2
		} else if role == "seller" {
			user.Role = 3
		}
		users = append(users, user)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users, nil

}

func (d *DB) ListAllLaptops() ([]entity.Laptop, error) {
	const op = "mysql.ListAllUsers"
	rows, err := d.conn.Connection().Query(`select laptops.id, laptops.cpu, laptops.ram, laptops.ssd, laptops.hdd, laptops.graphic,laptops.screen_size, laptops.company,laptops.price, laptops.image_url, laptops.redirect_url from laptops`)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)

	}
	defer rows.Close()

	var laptops []entity.Laptop

	// Iterate over the result set and populate the users slice
	for rows.Next() {
		var imageUrlTmp *string = new(string)
		var redirectUrlTmp *string = new(string)

		var laptop entity.Laptop
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

	return laptops, nil
}

func (d *DB) DeleteLaptop(LaptopID uint64) error {
	const op = "DeleteLaptop"
	// Execute the delete query

	_, err := d.conn.Connection().Query(`DELETE FROM user_laptop WHERE laptop_ref = ?`, LaptopID)
	if err != nil {
		return richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
	}

	_, err = d.conn.Connection().Query(`DELETE FROM laptops WHERE id = ?`, LaptopID)
	if err != nil {
		return richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
	}

	return nil
}

func (d *DB) UpdateLaptop(laptopID uint, updatedLaptop entity.Laptop) error {
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

	if !updatedLaptop.CreatedAt.IsZero() {
		updateQuery += " created_at = ?,"
		updateValues = append(updateValues, updatedLaptop.CreatedAt)
	}

	if updatedLaptop.ImageURL != "" {
		updateQuery += " image_url = ?,"
		updateValues = append(updateValues, updatedLaptop.ImageURL)
	}

	if updatedLaptop.RedirectURL != "" {
		updateQuery += " redirect_url = ?,"
		updateValues = append(updateValues, updatedLaptop.RedirectURL)
	}

	// Remove the trailing comma
	updateQuery = updateQuery[:len(updateQuery)-1]

	updateQuery += " WHERE id = ?"
	updateValues = append(updateValues, laptopID)

	// Execute the update query
	_, err := d.conn.Connection().Exec(updateQuery, updateValues...)
	return err
}

func (d *DB) RegisterAdmin() error {
	const op = "backofficeuser.RegisterAdmin"
	pass := []byte("admin12345")
	hashedPass, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
	}

	_, err = d.conn.Connection().Exec(`insert into users(name, phone_number, password, role) values (?, ? , ?, ?)`, "admin", "09111111111", hashedPass, entity.AdminRoleStr)
	if err != nil {
		return richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
	}

	return nil
}

func (d *DB) DeleteUser(UserID int) error {
	const op = "DeleteUser"
	// Execute the delete query

	_, err := d.conn.Connection().Query(`DELETE FROM user_laptop WHERE user_ref = ?`, UserID)
	if err != nil {
		return richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
	}

	_, err = d.conn.Connection().Query(`DELETE FROM users WHERE id = ?`, UserID)
	if err != nil {
		return richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
	}

	return nil
}
