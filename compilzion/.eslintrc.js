module.exports = {
  "env": {
    "node": true,
    "es2021": true,
    "jest/globals": true
  },
  "extends": [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended"
  ],
  "parser": "@typescript-eslint/parser",
  "parserOptions": {
    "ecmaVersion": 13,
    "sourceType": "module"
  },
  "plugins": [
    "@typescript-eslint",
    "jest"
  ],
  "rules": {
  }
};
