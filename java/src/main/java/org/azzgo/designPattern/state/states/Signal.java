package org.azzgo.designPattern.state.states;

import org.azzgo.designPattern.state.TrafficLight;

public interface Signal {
    String passable();
    void nextSignal(TrafficLight context);
}
