struct VecStack<T> {
    items: Vec<T>,
}


impl<T> VecStack<T> {
    pub fn pop(&mut self) -> Result<T, &str> {
        let item = self.items.pop();
        match item {
            Some(item) => Ok(item),
            None => Err("Empty stack can not pop!")
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
    pub fn new() -> VecStack<T>{
        return VecStack { items: Vec::new() }
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
}
