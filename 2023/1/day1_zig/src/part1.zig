const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;

test "input" {
    read_input();
}

pub fn read_input() void {
    const allocator = std.heap.page_allocator;
    const path = "input.txt";

    const contents = try readFile(allocator, path);
    defer allocator.free(contents);

    print("{}\n", .{contents});
}

fn readFile(allocator: std.mem.Allocator, path: []const u8) ![]u8 {
    const file = try std.fs.cwd().openFile(path, .{});
    defer file.close();

    const file_len = try file.getEndPos();
    const buffer = try allocator.alloc(u8, file_len);

    try file.readExactly(buffer);
    return buffer;
}
