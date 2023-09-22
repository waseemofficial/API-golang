#!/bin/sh

# Decrypt the file
mkdir $HOME/secrets
# --batch to prevent interactive command
# --yes to assume "yes" for questions
gpg --quiet --batch --yes --decrypt --passphrase="$Password" --output $HOME/secrets/my_secret.json my_secret.json.gpg

#export GOOGLE_APPLICATION_CREDENTIALS="$HOME/"
export NAMES="ji sunta ho"
echo -e "\e[32m $NAME \e[0m"