# randomWordChoice

## Usage

```
randomWordChoice [-t|--titlecase] [-s|--separator SEP] <wordlist1> <wordlist2>
```

A random word will be sampled from each word list, in the order provided.

By default, will return space-separated. Can provide your own separator, and also choose for each word to be capitalised.

Example:

``` sh
randomWordChoice $HOME/words/adjectives.txt $HOME/words/animals.txt
# => angry badger

randomWordChoice -t -s "-" $HOME/words/animals.txt $HOME/words/jobs.txt
# => Shrew-Postman
```

Originally created `animalhash` to provide something like docker's auto generated process names (using word lists _adjectives.txt_, _colours.txt_, and _animals.txt_). Then, decided to add in _jobs.txt_ as a way of generating some drawing prompts.

## Install

`go install`
