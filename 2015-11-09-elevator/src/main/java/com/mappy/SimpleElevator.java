package com.mappy;

import java.util.List;

import static com.google.common.collect.Lists.*;

public class SimpleElevator implements Elevator {
    private final List<String> commands = newArrayList(
            "OPEN",
            "CLOSE",
            "UP",
            "OPEN",
            "CLOSE",
            "UP",
            "OPEN",
            "CLOSE",
            "UP",
            "OPEN",
            "CLOSE",
            "UP",
            "OPEN",
            "CLOSE",
            "UP",
            "OPEN",
            "CLOSE",
            "DOWN",
            "DOWN",
            "DOWN",
            "DOWN",
            "DOWN");
    private int i = 0;

    @Override
    public void reset() {
        i = 0;
    }

    @Override
    public void userHasEntered() {}

    @Override
    public void userHasExited() {}

    @Override
    public void call(String atFloor, String to) {}

    @Override
    public void go(String floorToGo) {}

    @Override
    public String nextCommand() {
        return commands.get(i++ % commands.size());
    }
}
