# Efficient writing

limit number of syscal (write more data at once)

var w io.WriteCloser 
// initialise writer 
defer w.Close()
b := bufio.NewWriter(w) 
defer b.Flush()
// write operations