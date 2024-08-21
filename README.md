# Genpass

This is a CLI project develop in `Golang` primary for personal use.

It manages and creates passwords locally. You can optionally set up users to store those passwords.

If you use users, all information is store in a local `sqlite3` database under `$HOME/.genpass` folder.

Genpass by default encrypts the generated passwords and the entities for which the passwords are generated 
with the user password. Which in turn is saved in a hash form in a local database.

This projects is develop using the Cobra library. For more information visit [https://github.com/spf13/cobra](https://github.com/spf13/cobra).

To start first install the CLI app using:
```zsh
go install github.com/tizor98/genpass
``` 
And execute 
```zsh
genpass
``` 
or
```zsh
genpass help new
```
To get started.
