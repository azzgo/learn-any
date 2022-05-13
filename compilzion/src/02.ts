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
  constructor(private type: TokenType, private text: string) {}

  getType() {
    return this.type;
  }

  getValue() {
    return this.text;
  }

  toString() {
    return `${this.type}:    ${this.text}`;
  }
}

export function tokenize(text: string): Token[] {
  const chars = text.split("");
  let state = DfaState.Initial as DfaState;
  let tokenText = "";
  let tokenType: TokenType | null = null;
  const tokens: Token[] = [];

  const isAlpah = (char: string) => char.match(/^[A-z]$/);
  const isDigital = (char: string) => char.match(/^[0-9]$/);

  function initToken(char: string): DfaState {
    if (tokenText.length > 0 && tokenType !== null) {
      tokens.push(new Token(tokenType, tokenText));
      tokenText = "";
    }

    let newState = DfaState.Initial;

    if (isAlpah(char)) {
      newState = DfaState.Id;
      tokenType = TokenType.Identifier;
      tokenText += char;
    } else if (isDigital(char)) {
      newState = DfaState.IntLiteral;
      tokenType = TokenType.IntLiteral;
      tokenText += char;
    } else if (char === ">") {
      newState = DfaState.GT;
      tokenType = TokenType.GT;
      tokenText += char;
    }
    return newState;
  }

  for (const char of chars) {
    switch (state) {
      case DfaState.Initial:
        state = initToken(char);
        break;
      case DfaState.Id:
        if (isDigital(char) || isAlpah(char)) {
          tokenText += char;
        } else {
          state = initToken(char);
        }
        break;
      case DfaState.IntLiteral:
        if (isDigital(char)) {
          tokenText += char;
        } else {
          state = initToken(char);
        }
        break;
      case DfaState.GT:
        state = initToken(char);
        break;
      default:
        state = DfaState.Initial;
    }
  }

  // push token of last one  and reset to Initial state
  if (tokenText.length > 0) {
    initToken('');
  }
  return tokens;
}
