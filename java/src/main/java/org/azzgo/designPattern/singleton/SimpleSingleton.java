package org.azzgo.designPattern.singleton;

class SimpleSingleton {
    private static SimpleSingleton simpleSingleton;

    public static SimpleSingleton getInstance() {
        if (simpleSingleton == null) {
            simpleSingleton = new SimpleSingleton();
        }
        return simpleSingleton;
    }

    @Override
    public String toString() {
        return "SimpleSingleton";
    }
}
