#[derive(Debug)]
struct VecStack<T> {
    items: Vec<T>,
}

impl<T> VecStack<T> {
    pub fn pop(&mut self) -> Result<T, &str> {
        let item = self.items.pop();
        match item {
            Some(item) => Ok(item),
            None => Err("Empty stack can not pop!"),
        }
    }

    pub fn peek(&self) -> Result<&T, &str> {
        let item = self.items.last();
        match item {
            Some(item) => Ok(item),
            None => Err("Empty stack can not peek!"),
        }
    }
    pub fn push(&mut self, item: T) {
        self.items.push(item);
    }

    pub fn is_empty(&self) -> bool {
        return self.items.is_empty();
    }

    pub fn len(&self) -> usize {
        return self.items.len();
    }
}

impl<T> VecStack<T> {
    pub fn new() -> VecStack<T> {
        return VecStack { items: Vec::new() };
    }
}

#[cfg(test)]
mod tests {
    use super::VecStack;

    #[test]
    fn it_should_create_normal() {
        let mut stack = VecStack::<i32>::new();
        assert!(stack.is_empty());
        stack.push(3);
        stack.push(6);
        assert!(!stack.is_empty());
        assert_eq!(stack.len(), 2);
        let num = stack.pop();
        if let Ok(n) = num {
            assert_eq!(n, 6);
        }
    }

    fn check_parenthese(bracket_str: &str) -> bool {
        let mut match_bracket_stack = VecStack::new();
        let chars: Vec<&str> = bracket_str.split("").collect();
        for ch in chars {
            if ["(", "[", "{"].contains(&ch) {
                match_bracket_stack.push(ch);
            }

            if [")", "]", "}"].contains(&ch) {
                if match_bracket_stack.len() == 0 {
                    return false;
                }
                let item = match_bracket_stack.peek();
                if let Ok(ch1) = item {
                    if ch1 == &"(" && ch == ")" || ch1 == &"{" && ch == "}" || ch1 ==&"[" && ch == "]" {
                        let _ = match_bracket_stack.pop();
                    }
                }
            }
        }

        return match_bracket_stack.len() == 0;
    }

    #[test]
    fn it_used_to_match_brackets() {
        assert!(check_parenthese("[()]{}{[()()]()}"));
        assert!(!check_parenthese("[(])"));
        assert!(check_parenthese("[]{}[,]"));
    }


    #[test]
    fn it_final_stack_should_be_size1_and_it_item() {
        let mut stack = VecStack::new();
        stack.push("it");
        stack.push("was");
        let _ = stack.pop();
        stack.push("the");
        stack.push("best");
        let _ = stack.pop();
        stack.push("of");
        stack.push("times");
        let _ = stack.pop();
        let _ = stack.pop();
        let _ = stack.pop();
        stack.push("it");
        stack.push("was");
        let _ = stack.pop();
        stack.push("the");
        let _ = stack.pop();
        let _ = stack.pop();
        assert_eq!(stack.len(), 1);
        assert_eq!(&"it", stack.peek().unwrap());
    }
}
