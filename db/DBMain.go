package db

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const USERS_FILE = "users"

func AddNewUser(user int64) {
	users, err := os.Open(USERS_FILE + ".users")
	if err != nil {
		dbFileCreator()
		users, err = os.Open(USERS_FILE + ".users")
	}

	defer users.Close()

	_, err = users.WriteString(strconv.FormatInt(user, 10))
	if err != nil {
		log.Fatalf("File can't be writed; %s\n", err)
	}

	users.Sync()
}

func GetUsers() (users_ids []int64) {
	users, err := os.Open(USERS_FILE + ".users")
	if err != nil {
		log.Fatalf("File can't be opened; %s\n", err)
	}

	defer users.Close()

	one_user_id := bufio.NewScanner(users)

	for one_user_id.Scan() {
		id, _ := strconv.ParseInt(one_user_id.Text(), 10, 64)
		users_ids = append(users_ids, id)
	}

	err = one_user_id.Err()
	if err != nil {
		log.Fatal(err)
	}

	return users_ids
}

func dbFileCreator() {
	packer_script, err := os.Create(USERS_FILE + ".users")
	if err != nil {
		log.Fatalf("File can't be created; %s\n", err)
	}

	defer packer_script.Close()
	packer_script.Sync()
}
