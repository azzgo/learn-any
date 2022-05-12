import { tokenize, Token, TokenType } from './02';

it("should tokenize", () => {
  const result = tokenize('age >= 45');

  expect(result).toEqual([
    new Token(TokenType.Identifier, "age"),
    new Token(TokenType.GT, ">"),
    new Token(TokenType.IntLiteral, "45"),
  ]);
})
