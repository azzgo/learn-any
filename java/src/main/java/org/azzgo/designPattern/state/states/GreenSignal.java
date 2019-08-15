package org.azzgo.designPattern.state.states;

import org.azzgo.designPattern.state.TrafficLight;

public class GreenSignal implements Signal {
    @Override
    public void nextSignal(TrafficLight context) {
        context.setSignal(new YellowSignal());
    }

    @Override
    public String toString() {
        return "绿灯";
    }

    @Override
    public String passable() {
        return "阁下，当前绿灯，请通过";
    }
}
