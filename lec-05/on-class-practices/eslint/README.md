## Eslint

### Basic
- Config file:
> .eslintrc, .eslintrc.js

- Using for env
```json
{
  "env": {
    "browser": true
  }
}
```
- Parser options for ES6
```json
{
  "parserOptions": {
    "ecmaVersion": 6,
    "sourceType": "module",
    "ecmaFeatures": {
      "jsx": true
    }
  }
}
```
- Rules
```json
{
  "rules": {
    "eqeqeq": "error"
  }
}
```

## Resource
- [eslint](https://eslint.org/docs/7.0.0/user-guide/configuring)
- [eslint rules](https://eslint.org/docs/rules/)
- [airbnb-eslint](https://www.npmjs.com/package/eslint-config-airbnb)
