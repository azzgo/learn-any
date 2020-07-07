struct Stack<T> {
    items: Vec<T>,
}


impl<T> Stack<T> {
    pub fn pop(&mut self) -> Option<T> {
        let item = self.items.pop();
        return item;
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

impl<T> Stack<T> {
    pub fn new() -> Stack<T>{
        return Stack { items: Vec::new() }
    }
}

#[cfg(test)]
mod tests {
    use super::Stack;

    #[test]
    fn it_should_create_normal() {
        let mut stack = Stack::<i32>::new();
        assert!(stack.is_empty());
        stack.push(3);
        stack.push(6);
        assert!(!stack.is_empty());
        assert_eq!(stack.len(), 2);
        let num = stack.pop();
        if let Some(n) = num {
            assert_eq!(n, 6);
        }
    }
}
