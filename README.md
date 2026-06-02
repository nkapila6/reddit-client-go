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
    "title": "Maybe throw in \"racist\" while you're at it. What a joke.",
    "author": "Celatine_",
    "comments": "1",
    "permalink": "https://old.reddit.com/r/antiai/comments/1tv2ghe/maybe_throw_in_racist_while_youre_at_it_what_a/",
    "image": "https://i.redd.it/r87nu4i7ex4h1.png",
    "domain": "i.redd.it",
    "time": "Wednesday, June 3, 2026 at 12:03:32 AM"
  },
  {
    "subreddit": "antiai",
    "title": "Reminder that AI bros frequently utilize DARVO to defend AI",
    "author": "PhysicalBuy2566",
    "comments": "10",
    "permalink": "https://old.reddit.com/r/antiai/comments/1tv1jgz/reminder_that_ai_bros_frequently_utilize_darvo_to/",
    "image": "https://i.redd.it/0e2964ak8x4h1.jpeg",
    "domain": "i.redd.it",
    "time": "Tuesday, June 2, 2026 at 11:31:21 PM"
  },
  {
    "subreddit": "antiai",
    "title": "Yall is THIS what ai bros call art? 😭🙏",
    "author": "let_me-die_please",
    "comments": "9",
    "permalink": "https://old.reddit.com/r/antiai/comments/1tv0uu8/yall_is_this_what_ai_bros_call_art/",
    "image": "https://i.redd.it/oimfm3jn3x4h1.png",
    "domain": "i.redd.it",
    "time": "Tuesday, June 2, 2026 at 11:07:34 PM"
  },
  {
    "subreddit": "antiai",
    "title": "This is embarrassing.",
    "author": "Alert_Pie3002",
    "comments": "5",
    "permalink": "https://old.reddit.com/r/antiai/comments/1tuzxwg/this_is_embarrassing/",
    "image": "https://i.redd.it/d6jbxfgbyw4h1.png",
    "domain": "i.redd.it",
    "time": "Tuesday, June 2, 2026 at 10:36:18 PM"
  },
  {
    "subreddit": "antiai",
    "title": "Almonds also have protein and are delicious with chocolate.",
    "author": "c-k-q99903",
    "comments": "54",
    "permalink": "https://old.reddit.com/r/antiai/comments/1tuzu4p/almonds_also_have_protein_and_are_delicious_with/",
    "image": "https://i.redd.it/tleax234yw4h1.jpeg",
    "domain": "i.redd.it",
    "time": "Tuesday, June 2, 2026 at 10:32:46 PM"
  }
]
```
