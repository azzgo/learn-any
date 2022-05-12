// 02 | 正则文法和有限自动机：纯手工打造词法分析器

enum DfaState {
  Initial,
  Id,
  IntLiteral,
  GT,
}

export enum TokenType {
  Identifier,
  IntLiteral,
  GT,
}

export class Token {
  constructor(private type: TokenType, private value: string) {}

  getType() {
    return this.type;
  }

  getValue() {
    return this.value;
  }
}

export function tokenize(text: string): Token[] {
  return []
}

