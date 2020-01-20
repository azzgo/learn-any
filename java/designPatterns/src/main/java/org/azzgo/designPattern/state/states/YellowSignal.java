package org.azzgo.designPattern.state.states;

import org.azzgo.designPattern.state.TrafficLight;

public class YellowSignal implements Signal {

    @Override
    public void nextSignal(TrafficLight context) {
        context.setSignal(new RedSignal());
    }

    @Override
    public String toString() {
        return "黄灯";
    }

    @Override
    public String passable() {
        return "请速速通过。";
    }
}
