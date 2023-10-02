#!/bin/bash

# Prompt user for MariaDB username
read -p "Enter MariaDB username: " db_username

# Prompt user for MariaDB password (without showing the input)
read -s -p "Enter MariaDB password: " db_password
echo

# Prompt user for password (without showing the input)
read -s -p "Enter the new password of the admin: " password
echo

# Calculate the MD5 hash of the password
md5hash=$(echo -n "$password" | md5sum | awk '{print $1}')

# Construct the SQL query
sql_query="USE bims; UPDATE Users SET Password='$md5hash' WHERE ID = 1;"

# Run the SQL query using the mysql command-line tool
echo "$sql_query" | mysql -u "$db_username" -p"$db_password"

echo "SQL query executed successfully!"