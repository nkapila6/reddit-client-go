# reddit-client-go
reddit http client go

## example usage
```bash
go run main.go -sub antiai -sort new -limit 5 -images true| jq
```

## example output
```json
[
  {
    "subreddit": "antiai",
    "title": "This is embarrassing.",
    "permalink": "https://i.redd.it/d6jbxfgbyw4h1.png",
    "url": "https://i.redd.it/d6jbxfgbyw4h1.png",
    "image": "https://i.redd.it/d6jbxfgbyw4h1.png",
    "domain": "i.redd.it"
  },
  {
    "subreddit": "antiai",
    "title": "Almonds also have protein and are delicious with chocolate.",
    "permalink": "https://i.redd.it/tleax234yw4h1.jpeg",
    "url": "https://i.redd.it/tleax234yw4h1.jpeg",
    "image": "https://i.redd.it/tleax234yw4h1.jpeg",
    "domain": "i.redd.it"
  },
  {
    "subreddit": "antiai",
    "title": "the only AI I can get behind 🦅🦅🦅",
    "permalink": "https://i.redd.it/f1irqtssuw4h1.png",
    "url": "https://i.redd.it/f1irqtssuw4h1.png",
    "image": "https://i.redd.it/f1irqtssuw4h1.png",
    "domain": "i.redd.it"
  },
  {
    "subreddit": "antiai",
    "title": "My main question is how do the chicken heads fit?",
    "permalink": "https://i.redd.it/vslue5nupw4h1.png",
    "url": "https://i.redd.it/vslue5nupw4h1.png",
    "image": "https://i.redd.it/vslue5nupw4h1.png",
    "domain": "i.redd.it"
  },
  {
    "subreddit": "antiai",
    "title": "Martin Scorsese joins ai in film use",
    "permalink": "https://i.redd.it/o8lil9f2ow4h1.jpeg",
    "url": "https://i.redd.it/o8lil9f2ow4h1.jpeg",
    "image": "https://i.redd.it/o8lil9f2ow4h1.jpeg",
    "domain": "i.redd.it"
  }
]
```
