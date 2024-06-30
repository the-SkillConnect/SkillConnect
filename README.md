# SkillConnect
A Freelancing Telegram Mini App (TMA)

## Insert a user
```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation InsertUser { insertUser(input: { firstname: \"milad\", surname: \"rrr\", mobilePhone: \"5168416\", email: \"mili@gmail.com\", password: \"1234\" }) { email firstname id mobilePhone password surname } }"
  }' \
  http://localhost:8585/graphql
```

## Insert a project
```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation InsertProject { insertProject(input: { total_amount: \"123.1\", status: false, user_id: 1, fee: \"84.01\", title: \"tit\", description: \"dec\" }) { description fee id orderDate status title total_amount user_id } }"
  }' \
  http://localhost:8585/graphql

```