Anki - page - 75


When I executed script the via `ssh -i ~/.ssh/greenlight.pem -t ubuntu@52.90.171.119 "sudo bash /home/ubuntu/setup/01.sh"` looks like it didn't executed the following lines 


`sudo mkdir -p /home/${USERNAME}/.ssh`
`sudo cp -r /home/ubuntu/.ssh/authorized_keys /home/${USERNAME}/.ssh/`
`sudo chown -R ${USERNAME}:${USERNAME} /home/${USERNAME}/.ssh`
`sudo chmod 700 /home/${USERNAME}/.ssh`
`sudo chmod 600 /home/${USERNAME}/.ssh/authorized_keys`

When I ssh as a `ubuntu` user and executed line by line, I am able to do the ssh via new user

SSH hung up after internet discomment, how to end the session 
~.

Htop - for activity monitor 

ls -a for all files including hidden

mac terminal  - ctr a, ctr e

/root/.ssh/authorized_keys

# sed -i -e 's/.*exit 142" \(.*$\)/\1/' /root/.ssh/authorized_keys

/root/.ssh/authorized_keys contains a line `Please login as the user "ubuntu" rather than the user "root".` removing that allow to login via greenlight


https://repost.aws/knowledge-center/new-user-accounts-linux-instance


Copy Setup/01.sh

`rsync -e "ssh -i ~/.ssh/greenlight.pem" -rP --delete ./remote/setup ubuntu@3.93.241.94:/home/ubuntu/`

Execute the script

`ssh -i ~/.ssh/greenlight.pem -t ubuntu@3.93.241.94 "sudo bash /home/ubuntu/setup/01.sh"`

SSH
`ssh -i ~/.ssh/greenlight.pem -t ubuntu@3.93.241.94`





# Migration error
Getting error while running the migration 
- login as a greenlight and run the migration command directly  - `migrate -path ~/migrations -database $GREENLIGHT_DB_DSN up`

Log in as Postgres User
sudo -u postgres psql

Get the DB owner name 
`SELECT datname, pg_catalog.pg_get_userbyid(datdba) AS owner FROM pg_database WHERE datname = 'greenlight';`

Change the ownership of DB
`ALTER DATABASE greenlight OWNER TO greenlight;`

psql meta command 
- `\dt` table name
- `\l` database list


systemd - 
unit file - https://www.digitalocean.com/community/tutorials/understanding-systemd-units-and-unit-files
https://gist.github.com/ageis/f5595e59b1cddb1513d1b425a323db04
journalctl : https://www.digitalocean.com/community/tutorials/how-to-use-journalctl-to-view-and-manipulate-systemd-logs
