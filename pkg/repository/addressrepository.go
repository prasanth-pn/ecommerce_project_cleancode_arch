package repository

import (
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/domain"
	"errors"
	"fmt"
)

func (c *userDatabase) AddAddress(address domain.Address) error {
	query := `INSERT INTO addresses(user_id,f_name,l_name,phone_number,pincode,house,area,landmark,city)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9);`
	err := c.DB.QueryRow(query, address.User_Id, address.FName,
		address.LName, address.Phone_Number,
		address.Pincode, address.House, address.Area, address.Landmark, address.City).Err()
	return err
}

func (c *userDatabase) ListAddress(user_id uint) ([]domain.Address, error) {
	var address []domain.Address
	query := `SELECT address_id,user_id,f_name,l_name,phone_number,pincode,house,area,landmark,city
	FROM addresses
	where user_id=$1;`

	row, err := c.DB.Query(query, user_id)
	fmt.Println(err)
	if err != nil {
		return nil, errors.New("error happend while querying")
	}

	defer row.Close()
	for row.Next() {
		var addres domain.Address
		err = row.Scan(&addres.Address_id,
			&addres.User_Id,
			&addres.FName,
			&addres.LName,
			&addres.Phone_Number,
			&addres.Pincode,
			&addres.House,
			&addres.Area,
			&addres.Landmark,
			&addres.City)
		fmt.Println(addres.Area)

		if err != nil {
			return nil, errors.New("error while scaning response")
		}

		address = append(address, addres)
		fmt.Println(address)
	}
	return address, nil
}
