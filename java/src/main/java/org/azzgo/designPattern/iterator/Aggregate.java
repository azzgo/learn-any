package org.azzgo.designPattern.iterator;

import java.util.Iterator;

public interface Aggregate<T> {
    Iterator<T> iterator();
}
