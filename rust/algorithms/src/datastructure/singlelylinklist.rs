struct Node<T> {
    item: T,
    next: Option<Box<Node<T>>>,
}

struct SinglelyLinkList<T> {
    head: Option<Box<Node<T>>>,
    length: usize,
}

impl<T> SinglelyLinkList<T> {
    pub fn new() -> SinglelyLinkList<T> {
        SinglelyLinkList {
            head: None,
            length: 0,
        }
    }

    pub fn get(&mut self, k: usize) -> Option<&T> {
        if k > self.length {
            return None;
        }

        let mut cur = &self.head;
        let mut loop_index = 0;

        loop {
            if let Some(node) = cur {
                if loop_index == k {
                    return Some(&node.item);
                }

                cur = &node.next;
                loop_index += 1;
            } else {
                return None;
            }
        }
    }

    pub fn push(&mut self, item: T) {

    }

    pub fn delete(&mut self, k: usize) -> Option<T> {
        if k > self.length {
            return None;
        }

        let mut cur = &self.head;
        let mut prev: &Option<Box<Node<T>>> = &None;
        let mut loop_index = 0;

        loop {
            if let Some(node) = cur {
                if loop_index == k {
                    return match prev {
                        Some(prev) => {
                            prev.as_mut().next = node.next;
                            return Some(node.item);
                        },
                        None => None,
                    }
                }

                prev = cur;
                cur = &node.next;
                loop_index += 1;
            } else {
                return None;
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::SinglelyLinkList;

    #[test]
    #[cfg(feature = "option_expect_none")]
    fn it_should_be_none() {
        let link_list = SinglelyLinkList::<i32>::new();

        let get_node = link_list.get(3);
        let delete_node = link_list.delete(3);

        get_node.expect_none("期望返回 None");
        delete_node.expect_none("期望返回none")
    }
}
