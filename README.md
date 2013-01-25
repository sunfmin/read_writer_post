# Reader, Writer


## Simple Hello World

Let's say you have a file `/tmp/hello.txt` which has content of:

```
Hello World, 

This is my text for the Golang Reader.

```

Here you easy read the file with Golang


```
	var r io.Reader
	var err error
	r, err = os.Open("/tmp/hello.txt")
	if err != nil {
		panic(err)
	}
	var body []byte
	body, err = ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

```

source: http://play.golang.org/p/did0K6V9fv

Note that we used `ioutil.ReadAll` to read the full content of the text file. but this is not normally the way you do it.

## What's dangerous about ReadAll

Say now the file becomes `/tmp/25G_big_video_file.mp4`, the above program will read all the 25G content into the `body` bytes array, If you memory is not big enough to 25G, which is normally not nowadays, It will probably blow off your memory.

So don't do ReadAll unless you know the content is absolutely small.

## Consume stream chunk by chunk

Now let's suppose you have a list of text files, and You want to find out which files are containing the keyword `Golang`

You are attempting to do:

```
	body, _ = ioutil.ReadAll(r)
	found := strings.Index(string(body), "Golang")

```
source: http://play.golang.org/p/gUybiHGhFp

But which is simply a bad idea, Because the file might be huge, instead of doing that, We read the stream little by little.

Here is how we try to find the `Golang`, We read stream byte by byte, until we meet `G`, then we read the following 5 bytes to see if it's `olang`, If it is, then we found, If not until the end of file, means we can't find it.


```

func findGolang(r io.Reader) (position int) {
	var err error

	var G = []byte("G")
	var olang = []byte("olang")

	currentByte := make([]byte, 1)
	readByte := 0

	var l int
	for {
		l, err = r.Read(currentByte)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		readByte += l

		if bytes.Compare(currentByte, G) == 0 {
			next5 := make([]byte, 5)
			l, err = r.Read(next5)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			if bytes.Compare(next5, olang) == 0 {
				return readByte - 1
			}
		}
	}
	return -1
}

```

Yes, it's much longer implementation, But that's efficiency cost.



