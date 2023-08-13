package main

import (
	"bufio"
	"crypto/md5"
	"crypto/tls"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func main() {
	//read1()
	//write1()
	//write2()
	//write3()
	//mkdir1()
	//fileMode1()
	//fileOpen1()
	//fileOpen2()
	//stdin1()
	//newReader1()
	//newScanner1()
	//scanline2()
	//scanWords1()
	//scanWords2()
	//scanWords3()
	//pkgIo1()
	//pkgIo2()
	//pkgIo3()
	//pkgIo4()
	//pkgIo5()
	//pkgIo6()
	//pkgIo7()
	//pkgIOUtil()
	//pkgIoutil2()
	//pkgIoutil3()
	//encoding1()
	//encoding2()
	//pkgWalk()
	//pkgWalk2()
	//pkgWalk3()
	//pkgTime()
	//pkgTime1()
	//pkgTime2()
	//pkgHash1()
	//pkgHash2()
	//pkgFilepath()
	//pkgReviewExercises()
	//pkgReviewExercises2()
	//pkgReviewExercises3()
	//pkgReviewExercises4()
	//pkgReviewExercises5()
	//pkgBufioScanner()
	//pkgJson1()
	//pkgTls()
	//pkgRPCClient()
	pkgTls()
	pkgTLSClient()
}

func read1() {
	//entry := os.Args[1]
	//fmt.Println(entry) // go run xxx.go haha, 这里会打印haha

	src, err := os.Open("src.txt")
	if err != nil {
		//panic(err)
		fmt.Println("err happened: ", err)
	}
	defer src.Close()

	dst, err := os.Create("dst.txt")
	if err != nil {
		panic(err)
	}
	defer dst.Close()
	bs := make([]byte, 5)

	fmt.Println("=====", bs) // [0 0 0 0 0]
	src.Read(bs)
	dst.Write(bs)
}

func write1() {
	dst, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalf("error creating destination file:%v ", err)
	}
	defer dst.Close()

	dst.Write([]byte("Hello World."))
}
func write2() {
	dst, err := os.Create("hello.txt")
	if err != nil {
		log.Fatalln("error creating destination file: ", err.Error())
	}
	defer dst.Close()

	bs := []byte("Put some phrase here.")

	_, err = dst.Write(bs)
	if err != nil {
		log.Fatalln("error writing to file: ", err.Error())
	}
}
func write3() {
	dst, err := os.Create("/Users/yangqian/Downloads/hello.txt")
	if err != nil {
		log.Fatalln("error creating destination file: ", err.Error())
	}
	defer dst.Close()

	bs := []byte("Put some phrase here.")

	_, err = dst.Write(bs)
	if err != nil {
		log.Fatalln("error writing to file: ", err.Error())
	}
}

func mkdir1() {
	//err := os.Mkdir("/Users/yangqian/Downloads/somefolderthatdoesntexist", 0x777) // 权限？使用指定权限创建目录
	err := os.Mkdir("/Users/yangqian/Downloads/somefolderthatdoesntexist", os.ModePerm)
	if err != nil {
		log.Fatalln("my program broke on mkdir: ", err.Error())
	}
	f, err := os.Create("/Users/yangqian/Downloads/somefolderthatdoesntexist/hello.txt")
	if err != nil {
		log.Fatalln("my program broke", err.Error())
	}
	defer f.Close()

	str := "Put some phrase here."
	bs := []byte(str)

	_, err = f.Write(bs)
	if err != nil {
		log.Fatalln("error writing to file:  ", err.Error())
	}
}

func fileMode1() {
	fmt.Println(os.ModeDir)
	fmt.Printf("%p\n", os.ModeDir)
	fmt.Printf("%d\n", os.ModeDir)
	fmt.Println(os.ModeAppend)
	fmt.Println(os.ModeExclusive)
	fmt.Println(os.ModeTemporary)
	fmt.Println(os.ModeSymlink)
	fmt.Println(os.ModeDevice)
	fmt.Println(os.ModeNamedPipe)
	fmt.Println(os.ModeSocket)
	fmt.Println(os.ModeSetuid)
	fmt.Println(os.ModeCharDevice)
	fmt.Println(os.ModeSticky)
	fmt.Println(os.ModeType)
	fmt.Println(os.ModePerm)
}

func fileOpen1() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("my program broke", err.Error())
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("my program broke again")
	}
	str := string(bs)
	fmt.Println(str)
}

func fileOpen2() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("my program broken opening", err.Error())
	}
	defer f.Close()

	nf, err := os.Create("newFile.txt")
	if err != nil {
		log.Fatalln("my program broke creating: ", err.Error())
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("my program broke reading: ", err.Error())
	}

	_, err = nf.Write(bs)
	if err != nil {
		log.Fatalln("my program broke writing: ", err.Error())
	}
}

func stdout1() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("my program broke: ", err.Error())
	}
	defer f.Close()

	io.Copy(os.Stdout, f)
}
func stdin1() {
	rdr := strings.NewReader("test")
	io.Copy(os.Stdout, rdr)

	fmt.Println(
		strings.Contains("test", "es"),
		strings.Count("test", "t"),
		strings.HasPrefix("test", "te"),
		strings.HasSuffix("test", "st"),
		strings.Index("test", "e"),
		strings.Join([]string{"a", "b"}, "-"),
		strings.Repeat("a", 5),
		strings.Replace("aaaa", "a", "b", 2),

		strings.Split("a-b-c-d-e", "-"),
		strings.ToLower("TEST"),
		strings.ToUpper("test"),
		strings.TrimSpace("  test     "),
	)

	arr := []byte("test")
	fmt.Println(arr)

	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(str)
}
func newReader1() {
	dst, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalf("error creating destination file:%v", err)
	}
	defer dst.Close()

	rdr := strings.NewReader("hello world-")

	io.Copy(dst, rdr)
}
func newScanner1() {
	src, err := os.Open("initial.txt")
	if err != nil {
		log.Printf("error opening source file: %v", err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			fmt.Println(">>>", strings.ToUpper(line[0:1])+line[1:], "\n")
		}
		//fmt.Println(">>>", line)
	}
}
func scanline2() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
func scanWords1() {
	src, err := os.Open("hello.txt")
	if err != nil {
		log.Printf("error opening source file: %v", err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > 0 {
			fmt.Print(strings.ToUpper(word[0:1])+word[1:], " ")
		}
	}
	fmt.Println()
}
func scanWords2() {
	const input = "Now is the winter of out discontent, \nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}
	fmt.Printf("%d\n", count)
}
func scanWords3() {
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}
}
func pkgIo1() {
	rdr := strings.NewReader("test")
	io.Copy(os.Stdout, rdr)
}
func cp1(srcName, dstName string) error {
	src, err := os.Open(srcName)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return fmt.Errorf("error creating destination file: %v", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("error writing to destination file: %v", err)
	}
	return nil
}
func pkgIo2() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage: 04_io-copy <SRC> <DST>")
	}
	srcName := os.Args[1]
	dstName := os.Args[2]

	err := cp1(srcName, dstName)
	if err != nil {
		log.Fatalln(err)
	}
}
func pkgIo3() {
	dst, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalf("error creating destination file: %v", err)
	}
	defer dst.Close()

	rdr := strings.NewReader("hello world")
	io.Copy(dst, rdr)
}
func pkgIo4() {
	src, err := os.Open("src.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst1, err := os.Create("dst1.txt")
	if err != nil {
		panic(err)
	}
	defer dst1.Close()

	dst2, err := os.Create("dst2.txt")
	if err != nil {
		panic(err)
	}
	defer dst2.Close()

	rdr := io.TeeReader(src, dst1)
	rdr = io.TeeReader(rdr, os.Stdout)

	io.Copy(dst2, rdr)
}
func pkgIo5() {
	src, err := os.Open("src.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst1, err := os.Create("dst1.txt")
	if err != nil {
		panic(err)
	}
	defer dst1.Close()

	dst2, err := os.Create("dst2.txt")
	if err != nil {
		panic(err)
	}
	defer dst2.Close()

	dst3, err := os.Create("dst3.txt")
	if err != nil {
		panic(err)
	}
	defer dst3.Close()

	// TeeReader返回一个将其从r读取的数据写入w的Reader接口。所有通过该接口对r的读取都会执行对应的对w的写入。没有内部的缓冲：写入必须在读取完成前完成。写入时遇到的任何错误都会作为读取错误返回。
	rdr := io.TeeReader(src, dst1)
	rdr = io.TeeReader(rdr, os.Stdout)
	rdr = io.TeeReader(rdr, dst2)

	io.Copy(dst3, rdr)
}
func pkgIo6() {
	src, err := os.Open("src.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create("dst.txt")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	bs := make([]byte, 5)
	// io.ReadFul(r Reader, buf []byte)
	// 从 r 中精确的读取len(buf)字节数据填充进buf。函数返回写入的字节数和错误。
	// 只有没有读取到字节时才可能返回EOF；如果读取了有但不够时遇到了EOF，函数会返回ErrUnexpectedEOF。只有返回值err为nil时，返回值n才会等于len(buf)
	n, err := io.ReadFull(src, bs)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	dst.Write(bs)
}
func pkgIo7() {
	src, err := os.Open("src.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create("dst.txt")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	// io.LimitReader(r Reader, n int64) Reader 从r中读取n个字节后以EOF停止，返回值接口的底层为*LimitReader类型
	// type LimitedReader struct{R Reader N int64} 从R中读取数据，但限制读取的数据量最多为N字节，每次调用Read方法都会更新N以标记剩余可以读取的字节数
	rdr := io.LimitReader(src, 7)
	io.Copy(dst, rdr)
}
func pkgIo8() {
	f, err := os.Create("hello.txt")
	if err != nil {
		log.Fatalln("my program broke", err.Error())
	}
	defer f.Close()

	str := "Put some phrase here."
	// 方法1：
	//bs := []byte(str)
	//_, err = f.Write(bs)
	// 方法2：
	_, err = io.WriteString(f, str)
	if err != nil {
		log.Fatalln("error writing to file: ", err.Error())
	}
}
func pkgIOUtil() {
	myPhrase := "mmm, licorice"
	rdr := strings.NewReader(myPhrase)

	// ioutil.ReadAll(r io.Reader)([]byte,error)
	// 从r中读取数据直到EOF或遇到error，返回读取的数据和遇到的错误。成功的调用返回err为nil而非EOF。因为本函数定义为读取r直到EOF，它不会讲读取返回的EOF视为应报告的错误。
	bs, err := ioutil.ReadAll(rdr)
	if err != nil {
		log.Fatalln("my program broke again.")
	}

	str := string(bs)
	fmt.Println(str)
}
func pkgIoutil2() {
	err := ioutil.WriteFile("hello.txt", []byte("Hello world==="), 0777)
	if err != nil {
		panic("something went wrong")
	}
}
func pkgIoutil3() {
	myStr := "Hello Everybody"

	err := ioutil.WriteFile("hey.txt", []byte(myStr), 0777)
	if err != nil {
		log.Fatalln("something went wrong!", err.Error())
	}

	f, err := os.Open("hey.txt")
	if err != nil {
		log.Fatalln("couldn't read file!", err.Error())
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("Readall crashed!", err.Error())
	}

	fmt.Println(string(bs))
}
func encoding1() {
	f, err := os.Open("./state_table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// 2. parse a csv file
	csvReader := csv.NewReader(f)
	for true {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(record)
	}
}
func encoding2() {
	f, err := os.Open("./state_table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for rowCount := 0; ; rowCount++ {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		columns := make(map[string]int)
		if rowCount == 0 {
			for idx, column := range record {
				columns[column] = idx
			}
		}
		fmt.Println(columns)
		break
	}
}
func pkgWalk() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, info.Name(), info.Size(), info.Mode(), info.IsDir())
		return nil
	})
}
func pkgWalk2() {
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fmt.Println(info.Name())
		return nil
	})
}
func md5file(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := md5.New()
	io.Copy(h, f)
	return h.Sum(nil)
}
func pkgWalk3() {
	filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		bs := md5file(path)
		fmt.Printf("%s %x\n", path, bs)
		return nil
	})
}
func pkgTime() {
	fmt.Println(time.Now())
}
func pkgTime1() {
	timeAsString := "2012-01-22"
	//fmt.Println(time.Parse("2006-01-01", timeAsString))
	fmt.Println(time.Parse("2006-01-01_this-does-not-compile", timeAsString))

	timeAsString1 := "01/22/2012"
	fmt.Println(time.Parse("01/01_this-does-not-compile/2006", timeAsString1))
	//fmt.Println(time.Parse("01/01_this-does-not-compile/2006", timeAsString1))

	timeAsString2 := "2006-Jan-02"
	fmt.Println(time.Parse(timeAsString2, "2013-Feb-03"))
}
func pkgTime2() {
	from, to := os.Args[1], os.Args[2]
	fromTime, _ := time.Parse("2006-01-01_this-does-not-compile", from)
	toTime, _ := time.Parse("2006-01-01_this-does-not-compile", to)

	dur := toTime.Sub(fromTime)
	fmt.Println("elapsed days:", int(dur/(time.Hour*24)))
}
func pkgHash1() {
	fmt.Println(os.Args[0], os.Args) // 0 为执行文件
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := fnv.New64()
	io.Copy(h, f)
	fmt.Println("The sum is: ", h.Sum64())
}
func pkgHash2() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("couldn't open file", err.Error())
	}
	defer f.Close()

	h := md5.New()
	io.Copy(h, f)
	fmt.Printf("The hash (sum) is: %x\n", h.Sum(nil))

	f.Seek(0, 0)
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("read all didn't read well", err.Error())
	}
	fmt.Printf("The hash (sum) is: %x\n", md5.Sum(bs))
}
func pkgFilepath() {
	var counter int
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		counter++
		fmt.Println(path)
		return nil
	})
	fmt.Println(counter)
}
func getGravatarHash(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)
	h := md5.New()
	io.WriteString(h, email)
	finalBytes := h.Sum(nil)
	finalString := hex.EncodeToString(finalBytes)
	return finalString
}
func pkgReviewExercises() {
	var email string = os.Args[1]
	var gravatarHash = getGravatarHash(email)
	fmt.Println(`<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title></title>
</head>
<body>
<img src="http://www.gravatar.com/avatar/` + gravatarHash + `" alt="user picture"/>
</body>
</html>`)
}
func wordCount(str string) map[string]int {
	var words []string = strings.Fields(str)
	m := make(map[string]int)
	for _, v := range words {
		m[v]++
	}
	return m
}
func pkgReviewExercises2() {
	str := `Robert Cohn was once middleweight boxing champion of Princeton. Do not think that I am very much impressed by that as a boxing title, but it meant a lot to Cohn. He cared nothing for boxing, in fact he disliked it, but he learned it painfully and thoroughly to counteract the feeling of inferiority and shyness he had felt on being treated as a Jew at Princeton. There was a certain inner comfort in knowing he could knock down anybody who was snooty to him, although, being very shy and a thoroughly nice boy, he never fought except in the gym. He was Spider Kelly's star pupil. Spider Kelly taught all his young gentlemen to box like featherweights, no matter whether they weighed one hundred and five or two hundred and five pounds. But it seemed to fit Cohn. He was really very fast. He was so good that Spider promptly overmatched him and got his nose permanently flattened. This increased Cohn's distaste for boxing, but it gave him a certain satisfaction of some strange sort, and it certainly improved his nose. In his last year at Princeton he read too much and took to wearing spectacles. I never met any one of his class who remembered him. They did not even remember that he was middleweight boxing champion. I mistrust all frank and simple people, especially when their stories hold together, and I always had a suspicion that perhaps Robert Cohn had never been middleweight boxing champion, and that perhaps a horse had stepped on his face, or that maybe his mother had been frightened or seen something, or that he had, maybe, bumped into something as a young child, but I finally had somebody verify the story from Spider Kelly. Spider Kelly not only remembered Cohn. He had often wondered what had become of him. Robert Cohn was a member, through his father, of one of the richest Jewish families in New York, and through his mother of one of the oldest. At the military school where he prepped for Princeton, and played a very good end on the football team, no one had made him race-conscious. No one had ever made him feel he was a Jew, and hence any different from anybody else, until he went to Princeton. He was a nice boy, a friendly boy, and very shy, and it made him bitter. He took it out in boxing, and he came out of Princeton with painful self-consciousness and the flattened nose, and was married by the first girl who was nice to him. He was married five years, had three children, lost most of the fifty thousand dollars his father left him, the balance of the estate having gone to his mother, hardened into a rather unattractive mould under domestic unhappiness with a rich wife; and just when he had made up his mind to leave his wife she left him and went off with a miniature-painter. As he had been thinking for months about leaving his wife and had not done it because it would be too cruel to deprive her of himself, her departure was a very healthful shock. The divorce was arranged and Robert Cohn went out to the Coast. In California he fell among literary people and, as he still had a little of the fifty thousand left, in a short time he was backing a review of the Arts. The review commenced publication in Carmel, California, and finished in Provincetown, Massachusetts. By that time Cohn, who had been regarded purely as an angel, and whose name had appeared on the editorial page merely as a member of the advisory board, had become the sole editor. It was his money and he discovered he liked the authority of editing. He was sorry when the magazine became too expensive and he had to give it up. By that time, though, he had other things to worry about. He had been taken in hand by a lady who hoped to rise with the magazine. She was very forceful, and Cohn never had a chance of not being taken in hand. Also he was sure that he loved her. When this lady saw that the magazine was not going to rise, she became a little disgusted with Cohn and decided that she might as well get what there was to get while there was still something available, so http://www.en8848.com.cn she urged that they go to Europe, where Cohn could write. They came to Europe, where the lady had been educated, and stayed three years. During these three years, the first spent in travel, the last two in Paris, Robert Cohn had two friends, Braddocks and myself. Braddocks was his literary friend. I was his tennis friend. The lady who had him, her name was Frances, found toward the end of the second year that her looks were going, and her attitude toward Robert changed from one of careless possession and exploitation to the absolute determination that he should marry her. During this time Robert's mother had settled an allowance on him, about three hundred dollars a month. During two years and a half I do not believe that Robert Cohn looked at another woman. He was fairly happy, except that, like many people living in Europe, he would rather have been in America, and he had discovered writing. He wrote a novel, and it was not really such a bad novel as the critics later called it, although it was a very poor novel. He read many books, played bridge, played tennis, and boxed at a local gymnasium. I first became aware of his lady's attitude toward him one night after the three of us had dined together. We had dined at l'Avenue's and afterward went to the Café de Versailles for coffee. We had several _fines_ after the coffee, and I said I must be going. Cohn had been talking about the two of us going off somewhere on a weekend trip. He wanted to get out of town and get in a good walk. I suggested we fly to Strasbourg and walk up to Saint Odile, or somewhere or other in Alsace. "I know a girl in Strasbourg who can show us the town," I said. Somebody kicked me under the table. I thought it was accidental and went on: "She's been there two years and knows everything there is to know about the town. She's a swell girl." I was kicked again under the table and, looking, saw Frances, Robert's lady, her chin lifting and her face hardening. "Hell," I said, "why go to Strasbourg? We could go up to Bruges, or to the Ardennes." Cohn looked relieved. I was not kicked again. I said good-night and went out. Cohn said he wanted to buy a paper and would walk to the corner with me. "For God's sake," he said, "why did you say that about that girl in Strasbourg for? Didn't you see Frances?" "No, why should I? If I know an American girl that lives in Strasbourg what the hell is it to Frances?" "It doesn't make any difference. Any girl. I couldn't go, that would be all." "Don't be silly." "You don't know Frances. Any girl at all. Didn't you see the way she looked?" "Oh, well," I said, "let's go to Senlis." "Don't get sore." "I'm not sore. Senlis is a good place and we can stay at the Grand Cerf and take a hike in the woods and come home." "Good, that will be fine." "Well, I'll see you to-morrow at the courts," I said. "Good-night, Jake," he said, and started back to the café. "You forgot to get your paper," I said. "That's so." He walked with me up to the kiosque at the corner. "You are not sore, are you, Jake?" He turned with the paper in his hand. "No, why should I be?" "See you at tennis," he said. I watched him walk back to the café holding his paper. I rather liked him and evidently she led him quite a life.`
	fmt.Println(wordCount(str))
}
func pkgReviewExercises3() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println(x)
	sort.Ints(x)
	fmt.Println(x)

	x = x[1:]
	fmt.Println(x)

	x = x[:(len(x) - 1)]
	fmt.Println(x)

	total := 0
	for _, value := range x {
		total += value
	}

	fmt.Println(total)
	fmt.Println(total / len(x))
}
func swap(x, y *int) {
	*x, *y = *y, *x
}
func pkgReviewExercises4() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println("x", x)
	fmt.Println("y", y)
}
func countClumps(xs []int) int {
	clumps := 0
	inClumps := false
	for i := 1; i < len(xs); i++ {
		curr, prev := xs[i], xs[i-1]
		if !inClumps && curr == prev {
			inClumps = true
			clumps++
		}
		if inClumps && curr != prev {
			inClumps = false
		}
	}
	return clumps
}
func pkgReviewExercises5() {
	clumps := countClumps([]int{1, 1, 1, 1, 1})
	fmt.Println(clumps)
}
func pkgReviewExercises6() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("my program broke", err.Error())
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("my program broke again.")
	}

	str := string(bs)
	fmt.Println(str)
}
func pkgReviewExercise() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("my program broke opening: ", err.Error())
	}
	defer f.Close()

	nf, err := os.Create("newFile.txt")
	if err != nil {
		log.Fatalln("my program broken creating: ", err.Error())
	}

	bc, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("my program broken reading: ", err.Error())
	}

	_, err = nf.Write(bc)
	if err != nil {
		log.Fatalln("my program broke writing: ", err.Error())
	}
}

func pkgReviewExercisesCP() {
	srcName := os.Args[1]
	dstName := os.Args[2]

	src, err := os.Open(srcName)
	if err != nil {
		log.Fatalln("my program broke opening: ", err.Error())
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		log.Fatalln("my program broke creating: ", err.Error())
	}
	defer dst.Close()

	bs, err := ioutil.ReadAll(src)
	if err != nil {
		log.Fatalln("my program broke reading: ", err.Error())
	}

	_, err = dst.Write(bs)
	if err != nil {
		log.Fatalln("my program broke writing: ", err.Error())
	}
}
func pkgReviewExercses() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage: o1_this-does-not-compile <SRC> <DST>")
	}
	srcName := os.Args[1]
	dstName := os.Args[2]

	src, err := os.Open(srcName)
	if err != nil {
		log.Fatalln("my program broke opening: ", err.Error())
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		log.Fatalln("my program broke creating: ", err.Error())
	}
	defer dst.Close()

	bs, err := ioutil.ReadAll(src)
	if err != nil {
		log.Fatalln("my program broke reading: ", err.Error())
	}

	_, err = dst.Write(bs)
	if err != nil {
		log.Fatalln("my program broke writing: ", err.Error())
	}
}

func packageIOCopy(srcName, dstName string) error {
	src, err := os.Open(srcName)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return fmt.Errorf("error creating destination file: %v", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("error writing to destination file: %v", err)
	}
	return nil
}

func pkgOsWriteSliceBytes() {
	dst, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalf("error creating destination file: %v", err)
	}
	defer dst.Close()

	dst.Write([]byte("Hello World."))
}

func pkgIoCopyStringNewReader() {
	dst, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalf("error creating destination file: %v", err)
	}
	defer dst.Close()

	rdr := strings.NewReader("hello world")

	io.Copy(dst, rdr)
}

func pkgBufioScanner() {
	src, err := os.Open("newFile.txt")
	if err != nil {
		log.Printf("error opening source file: %v", err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(">>>", line)
	}
}
func pkgJson1() {
	jsonData := `{
	"name": "Todd McLead"
	}`
	var obj map[string]string
	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
	fmt.Println(obj["name"])
	fmt.Printf("%T\n", obj["name"])
	//fmt.Println([]byte(jsonData))

	jsonData1 := `{
	"name": "Todd McLead",
	"age": 44
	}`
	var obj1 map[string]interface{}
	err = json.Unmarshal([]byte(jsonData1), &obj1)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj1)
	fmt.Println(obj1["name"])
	fmt.Println(obj1["age"])
	fmt.Printf("%T\n", obj1["name"])
	fmt.Printf("%T\n", obj1["age"])

	type Anything interface{}
	var obj2 map[string]Anything
	err = json.Unmarshal([]byte(jsonData1), &obj2)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj2)

	jsonData3 := `[100, 200, 300.5, 400.1234]`
	var obj3 []float64

	err = json.Unmarshal([]byte(jsonData3), &obj3)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj3)
	fmt.Printf("%T\n", obj3)
	fmt.Println(obj3[1])
	fmt.Printf("%T\n", obj3[1])

	jsonData4 := `100`
	var obj4 interface{}

	err = json.Unmarshal([]byte(jsonData4), &obj4)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj4)
	fmt.Printf("%T", obj4)

	jsondata5 := `100`
	var obj5 interface{}
	err = json.Unmarshal([]byte(jsondata5), &obj5)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj5)

	fmt.Printf("%T\n", obj5)
	x := 100 + obj5.(float64)
	fmt.Println(x)

	jsonData6 := `100`
	var obj6 interface{}
	err = json.Unmarshal([]byte(jsonData6), &obj6)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj6)
	fmt.Printf("%T\n", obj6)

	v, ok := obj6.(float64)
	if !ok {
		v = 0
	}
	y := 100 + v
	fmt.Println(y)

	jsondata7 := `100`
	var obj7 interface{}
	err = json.Unmarshal([]byte(jsondata7), &obj7)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj7)

	v7, ok7 := obj7.(float64)
	if !ok7 {
		v7 = 0
	}
	x7 := 100 + v7

	fmt.Println(x7)
	bs7, err := json.Marshal(x7) // []byte, err
	fmt.Println(string(bs7), err, bs7)

	jsondata8 := `100`
	var obj8 interface{}
	err = json.Unmarshal([]byte(jsondata8), &obj8)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj8)

	fmt.Println(obj8)
	fmt.Printf("%T\n", obj8)

	v8, ok8 := obj8.(float64)
	if !ok8 {
		v8 = 0
	}
	x8 := 100 + v8
	fmt.Println(x8)
	bs8, err := json.Marshal([]int{1, 2, 3, 4})
	fmt.Println(string(bs8), err)
}

// TCP SSl
type Result struct {
	Num, Ans int
}
type Cal int

func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}
func pkgRPC() {
	rpc.Register(new(Cal))
	rpc.HandleHTTP()

	log.Printf("Serving RPC server on port %d", 1234)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error serving: ", err)
	}
}
func pkgRPCClient() {
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	var result Result
	if err := client.Call("Cal.Square", 12, &result); err != nil {
		log.Fatal("Failed to call Cal.Square. ", err)

	}
	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
func pkgTls() {
	rpc.Register(new(Cal))
	cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	listener, _ := tls.Listen("tcp", ":1234", config)
	log.Printf("serving RPC server on port %d", 1234)
	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
}
func pkgTLSClient() {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, _ := tls.Dial("tcp", "localhost:1234", config)
	defer conn.Close()
	client := rpc.NewClient(conn)

	var result Result
	if err := client.Call("Cal.Square", 12, &result); err != nil {
		log.Fatal("Failed to call Cal.Square. ", err)
	}
	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
