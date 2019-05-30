package org.azzgo.designPattern.iterator;

import org.junit.Test;

import java.util.Iterator;

public class IteratorTest {
    @Test() public void play() {
        BookShelf bookShelf = new BookShelf();
        bookShelf.appendBook(new Book("环游世界80天"));
        bookShelf.appendBook(new Book("圣经"));
        bookShelf.appendBook(new Book("爱丽丝梦游仙境"));

        Iterator<Book> it = bookShelf.iterator();

        while (it.hasNext()) {
            Book book = it.next();
            System.out.println(book.getName());
        }
    }
}
