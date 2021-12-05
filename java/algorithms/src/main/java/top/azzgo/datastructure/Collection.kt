package top.azzgo.datastructure


class Bag<Item> : Iterable<Item> {
    constructor() {}

    fun add(item: Item) {

    }

    fun isEmpty(): Boolean {
        return false
    }

    fun size(): Int {
        return 0;
    }

    override fun iterator(): Iterator<Item> {
        TODO("Not yet implemented")
    }

}

class Quene<Item> : Iterable<Item> {
    private var first: LinkNode<Item>? = null
    private var last: LinkNode<Item>? = null
    private var size = 0


    constructor() {}

    fun enqueue(item: Item) {
        val oldLast = last
        last = LinkNode()
        last!!.item = item
        last!!.next = null
        // when only one item ,first = last
        if (isEmpty()) first = last
        else oldLast!!.next = last
        size++
    }

    fun dequeue(): Item? {
        if (isEmpty()) {
            return null
        }
        val item = first!!.item
        first = first!!.next
        // when only one item pop, the last point to the one, and first is empty, need point last to empty too
        if (isEmpty()) last = null
        size--
        return item
    }

    fun isEmpty(): Boolean {
        return first == null
    }

    fun size(): Int {
        return size
    }

    override fun iterator(): Iterator<Item> {
        TODO("Not yet implemented")
    }
}

interface Stack<Item> : Iterable<Item> {
    fun push(item: Item)
    fun pop(): Item?
    fun isEmpty(): Boolean
    fun size(): Int
}


class LinkStack<Item> : Stack<Item> {
    private var first: LinkNode<Item>? = null
    private var size = 0;

    inner class ListIterator : Iterator<Item> {
        private var current = first;
        override fun hasNext(): Boolean {
            return current != null
        }

        override fun next(): Item {
            val item = current!!.item;
            current = current!!.next;
            return item!!;
        }
    }

    override fun push(item: Item) {
        val oldNode = first
        first = LinkNode()
        first!!.item = item
        first!!.next = oldNode
        size++;
    }

    override fun pop(): Item? {
        if (first == null) {
            return null
        }
        val item = first!!.item
        first = first!!.next
        size--
        return item;
    }

    override fun isEmpty(): Boolean {
        return first == null
    }

    override fun size(): Int {
        return size
    }


    override fun iterator(): Iterator<Item> {
        return ListIterator()
    }
}

class ResizingStack<Item> : Stack<Item> {
    private var entries: MutableList<Item?> = MutableList(10) { null }
    private var size = 0

    inner class ReserveArrayIterator : Iterator<Item> {
        private var i = size

        override
        fun hasNext(): Boolean {
            return i > 0
        }

        override fun next(): Item {
            return entries[--i]!!;
        }

    }

    private fun resize(cap: Int) {
        val tempEntries = MutableList<Item?>(cap) { null }
        for ((index) in entries.withIndex()) {
            if (index < cap) {
                tempEntries[index] = entries[index];
            }
        }
        entries = tempEntries;
    }

    override fun push(item: Item) {
        if (size === entries.size) resize(2 * entries.size)
        entries[size++] = item;
    }

    override fun pop(): Item? {
        if (size == 0) {
            return null
        }
        val item = entries[--size]!!
        entries[size] = null;
        if (size > 0 && size === entries.size / 4) resize(entries.size / 2)
        return item;
    }

    override fun isEmpty(): Boolean {
        return size === 0
    }

    override fun size(): Int {
        return size;
    }

    override fun iterator(): Iterator<Item> {
        return ReserveArrayIterator()
    }

}
