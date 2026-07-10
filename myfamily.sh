#curl https://learn.reboot01.com/assets/superhero/all.json | jq ".[] | select(.id==1)| .connections.relatives"| sed 's/"//g'

curl -s https://learn.reboot01.com/assets/superhero/all.json | jq '.[] | select(.id=='$HERO_ID')| .connections.relatives'| tr -d '"'