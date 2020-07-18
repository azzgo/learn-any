struct Node<T> {
    item: T,
    next: Option<Box<Node<T>>>,
}

