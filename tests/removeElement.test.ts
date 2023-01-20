import { removeElement } from "../removeElement";



describe('Remove Element', () => {
  test.each([
    {
      name: "test 1",
      args: {
        nums: [3, 2, 2, 3],
        val: 3,
      },
      want: [2, 2],
    },
    {
      name: "test 2",
      args: {
        nums: [0, 1, 2, 2, 3, 0, 4, 2],
        val: 2,
      },
      want: [0, 1, 4, 0, 3],
    }
  ])(`$name`, ({ name, args, want }) => {
    expect(removeElement(args.nums, args.val)).toBe(want.length);
    expect(args.nums.slice(0, want.length).sort()).toEqual(want.sort());
  });
});
