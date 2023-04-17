# theychat

## 5 operation
1. SendMsg = 0
2. AddSingleChat = 1
3. AddGroupChat = 2
4. JoinSingleChat = 3
5. JoinGroupChat = 4

<details>
<summary>config.toml</summary>

**放在根目录下**

```
[detail]
page_size = 5
single_chat_cap = 2
group_chat_cap = 500
[redis]
host = "127.0.0.1"
port = "6379"
db = 0
[mysql]
user = "debian-sys-maint"
password = "9nVW7LM3BXm7mFZ9"
host = "localhost"
port = "3306"
dns = "debian-sys-maint:9nVW7LM3BXm7mFZ9@tcp(localhost:3306)/theychat?charset=utf8mb4&parseTime=True&loc=Local"
[token]
secret_key = "cgynbnbznb"
effect_time = 7200000000000
```

</details>