import { strStr } from "../strStr";



describe('Find the Index of the First Occurrence in a String', () => {
  test.each([
    {
      name: "test 1",
      args: {
        haystack: "sadbutsad", needle: "sad"
      },
      want: 0,
    },
    {
      name: "test 2",
      args: {
        haystack: "leetcode", needle: "leeto"
      },
      want: -1,
    }
  ])(`$name`, ({ name, args, want }) => {
    expect(strStr(args.haystack, args.needle)).toBe(want);
  });
});
