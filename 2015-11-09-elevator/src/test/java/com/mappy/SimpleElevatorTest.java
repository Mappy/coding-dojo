package com.mappy;

import java.util.List;

import org.junit.Test;

import static com.google.common.collect.Lists.*;
import static org.assertj.core.api.Assertions.*;

public class SimpleElevatorTest {
    private final Elevator elevator = new SimpleElevator();

    @Test
    public void should_do_omnibus() {
        for (int i = 0; i < 5; i++) {
            assertCommands(elevator, newArrayList("OPEN", "CLOSE", "UP"));
        }
        assertCommands(elevator, newArrayList("OPEN", "CLOSE", "DOWN", "DOWN", "DOWN", "DOWN", "DOWN", "OPEN"));
    }

    public void assertCommands(Elevator elevator, List<String> commands) {
        List<String> answers = newArrayList();
        for (int i = 0; i < commands.size(); i++) {
            answers.add(elevator.nextCommand());
        }
        assertThat(answers).containsExactlyElementsOf(commands);
    }
}
