package postgres

import (
	"fmt"
	"os"
	"rsweb/src/model"
	"rsweb/src/util"
	"strings"
)

func InsertUser(data model.TokenResponse) {
	family := strings.Split(os.Getenv("FAMILY"), ",")

	idx := util.Contains(family, data.Name)

	var role int
	if idx < 0 {
		role = 0
	} else {
		role = 10 + idx
	}

	sql := fmt.Sprintf("INSERT INTO public.tb_user (user_role, user_name, user_email, user_naverid) VALUES (%d, '%s', '%s', '%s')",
		role, data.Name, data.Email, data.Id)
	result, err := PostgresDB.Exec(sql)
	if err != nil {
		panic(err)
	}

	inserted, err := result.RowsAffected()
	if err != nil {
		panic(err)
	} else if inserted == 1 {
		fmt.Println("1 user successfully inserted")
	}
}

func SelectFamily() ([]model.Family, error) {
	sql := "SELECT user_role, user_naverid FROM public.tb_user WHERE user_role > 9"
	rows, err := PostgresDB.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	families := make([]model.Family, 0)
	for rows.Next() {
		var family model.Family
		err := rows.Scan(&family.Role, &family.Id)
		if err != nil {
			return nil, err
		}
		families = append(families, family)
	}

	return families, nil
}

func FindUserByNaverId(naverid string) int {
	sql := fmt.Sprintf("SELECT user_id FROM public.tb_user WHERE user_naverid='%s'", naverid)
	rows, err := PostgresDB.Query(sql)
	if err != nil {
		return -99
	}

	defer rows.Close()

	if rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return -99
		}
		return id
	}

	return -1
}

func FindRoleByNaverId(naverid string) int {
	sql := fmt.Sprintf("SELECT user_role FROM public.tb_user WHERE user_naverid='%s'", naverid)
	rows, err := PostgresDB.Query(sql)
	if err != nil {
		return -99
	}

	defer rows.Close()

	if rows.Next() {
		var role int
		err := rows.Scan(&role)
		if err != nil {
			return -99
		}
		return role
	}

	return -1
}
