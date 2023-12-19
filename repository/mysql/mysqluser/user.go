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

	res, err := d.conn.Connection().Exec(`insert into laptops(cpu, ram,ssd, hdd, grapic_card, screen_size, company, image_url, redirect_url ) values (?, ?, ?, ?, ?, ?, ?, ?, ?)`, laptop.CPU, laptop.RAM, laptop.SSD, laptop.HDD, laptop.GraphicCard, laptop.ScreenSize, laptop.Company, laptop.ImageURL, laptop.RedirectURL)
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

func (d *DB) GetLaptops(UserID uint) ([]entity.Laptop, error) {
	const op = "mysql.GetLaptops"
	var laptops []entity.Laptop
	rows, err := d.conn.Connection().Query(`select laptops.id, laptops.cpu, laptops.ram, laptops.ssd, laptops.hdd, laptops.grapic_card,laptops.screen_size, laptops.company, laptops.image_url, laptops.redirect_url from laptops join user_laptop on laptops.id = user_laptop.laptop_ref where user_laptop.user_ref = ?`, UserID)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	for rows.Next() {
		var laptop entity.Laptop
		err := rows.Scan(&laptop.ID, &laptop.CPU, &laptop.RAM, &laptop.SSD, &laptop.HDD, &laptop.GraphicCard, &laptop.ScreenSize, &laptop.Company, &laptop.ImageURL, &laptop.RedirectURL)
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
