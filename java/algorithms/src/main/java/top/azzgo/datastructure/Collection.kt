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

    constructor() {}

    fun enqueue(item: Item) {}
    fun dequeue(): Item {
        TODO("Not yet implemented")
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

interface Stack<Item> : Iterable<Item> {
    fun push(item: Item)
    fun pop(): Item
    fun isEmpty(): Boolean
    fun size(): Int
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

    override fun pop(): Item {
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
