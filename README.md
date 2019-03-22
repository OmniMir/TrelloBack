# Go Go TrelloBack

If you need local backup of your important boards, lists and cards, you use TrelloBack (not yet!!!).

One binary to rulу all yours [trello-kanban-lists-etc...](https://trello.com)



## Installation

Simply download one execution file from releases.



## Configuration

Because Trello use some limitations to 3rd-party API-request you need get API KEYS: 

Rename ```config.example.json``` to ```config.json```.

Log In or Sign Up to [Trello](https://trello.com), go https://trello.com/app-key and copy your key to ```api_key```.

Then you need generate API Token on the same page and copy your token to ```api_token```.



And last but not least -  input once and future folder of backups to ```destination_folder```. Remember use ```//``` instead of ```/```!!!



##  Use

Start TrelloBack execution file in GUI or CLI. Wait some time, make tea, do your deeds.

And you get as many json files as your **Trello Cards** in folders named by **Trello Organization** and **Trello Board**.

(Psst, organization name "Personal" reserved for boards without organization, sorry)

Each filename begins with name of **Trello List** and ends with name of **Trello Card**.

In json file you can find **your** valuable data that stores on **your** harddrive.

![](<https://raw.githubusercontent.com/OmniMir/TrelloBack/master/EXAMPLE.png>)

## In future

Incremental backup

More information of cards

Backuping to Zip

