import java.io.IOException;
import java.io.InputStream;
import java.nio.charset.StandardCharsets;
import java.util.Objects;
import java.util.function.Function;

public class Day12 {

	public static void main(String[] args) {
		Day12 day = new Day12(); // https://adventofcode.com/2024/day/12

		String full = readFile("input");

		day.run(full, day::part2, "Part 2 result");
	}

	public static String readFile(String fileName) {
		try (InputStream is = Day12.class.getResourceAsStream(fileName)) {
			Objects.requireNonNull(is, () -> "File %s not found".formatted(fileName));
			return new String(is.readAllBytes(), StandardCharsets.UTF_8);
		} catch (IOException e) {
			throw new RuntimeException(e);
		}
	}

	long run() {
		return run(readFile("input"));
	}

	long run(String input) {
		System.out.printf("Running %s%n", getClass().getSimpleName());
		long time2 = run(input, this::part2, "Part 2 result");
		return time2;
	}

	long run(String input, Function<String, String> function, String label) {
		long start = System.currentTimeMillis();
		String res = function.apply(input);
		long time = System.currentTimeMillis() - start;
		System.out.printf("[%d ms] %s: %s%n", time, label, res);
		return time;
	}

	private static final int AREA = 0;
	private static final int PERIMETER = 1;
	private static final int SIDES = 2;

	public String part2(String input) {
		char[][] map = parseMap(input);

		boolean[][] visited = new boolean[map.length][map[0].length];
		int result = 0;
		int[] region;
		for (int y = 0; y < map.length; y++) {
			for (int x = 0; x < map[0].length; x++) {
				if ((region = measureRegion(map, x, y, visited)) != null) {
					result += region[AREA] * region[SIDES];
				}
			}
		}

		return String.valueOf(result);
	}

	private static char[][] parseMap(String input) {
		return input.lines().map(String::toCharArray).toArray(char[][]::new);
	}

	private int[] measureRegion(char[][] map, int x, int y, boolean[][] visited) {
		if (visited[y][x]) {
			return null;
		}
		int[] region = new int[3];
		measureRegion(map, x, y, visited, map[y][x], region);
		return region;
	}

	private void measureRegion(char[][] map, int x, int y, boolean[][] visited, char plant, int[] region) {
		if (visited[y][x]) {
			return;
		}
		region[AREA]++;
		visited[y][x] = true;

		// up
		if (y > 0 && map[y - 1][x] == plant) {
			measureRegion(map, x, y - 1, visited, plant, region);
		} else {
			region[PERIMETER]++;
			boolean sideContinuation = x > 0 && map[y][x - 1] == plant && (y == 0 || map[y - 1][x - 1] != plant);
			if (!sideContinuation) {
				region[SIDES]++;
			}
		}

		// right
		if (x < map[0].length - 1 && map[y][x + 1] == plant) {
			measureRegion(map, x + 1, y, visited, plant, region);
		} else {
			region[PERIMETER]++;
			boolean sideContinuation = y > 0 && map[y - 1][x] == plant
					&& (x == map[0].length - 1 || map[y - 1][x + 1] != plant);
			if (!sideContinuation) {
				region[SIDES]++;
			}
		}

		// down
		if (y < map.length - 1 && map[y + 1][x] == plant) {
			measureRegion(map, x, y + 1, visited, plant, region);
		} else {
			region[PERIMETER]++;
			boolean sideContinuation = x < map[0].length - 1 && map[y][x + 1] == plant
					&& (y == map.length - 1 || map[y + 1][x + 1] != plant);
			if (!sideContinuation) {
				region[SIDES]++;
			}
		}

		// left
		if (x > 0 && map[y][x - 1] == plant) {
			measureRegion(map, x - 1, y, visited, plant, region);
		} else {
			region[PERIMETER]++;
			boolean sideContinuation = y < map.length - 1 && map[y + 1][x] == plant && (x == 0 || map[y + 1][x - 1] != plant);
			if (!sideContinuation) {
				region[SIDES]++;
			}
		}
	}

}
