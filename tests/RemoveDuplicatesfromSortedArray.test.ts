import { removeDuplicates } from "../RemoveDuplicatesfromSortedArray";


describe('remove duplicates from sorted array', () => {
  test.each([
    {
      name: "test 1",
      args: {
        input: [1, 1, 2],
      },
      want: [1, 2],
    },
    {
      name: "test 2",
      args: {
        input: [0, 0, 1, 1, 1, 2, 2, 3, 3, 4],
      },
      want: [0, 1, 2, 3, 4],
    }
  ])(`$name`, ({ name, args, want }) => {
    expect(removeDuplicates(args.input)).toBe(want.length);
    expect(args.input.slice(0, want.length)).toEqual(want);
  });
});
