### **Experiment 3: Using Channels for Goroutine Communication**  

#### **Understanding Channels in Go**  
So far, we’ve seen how to launch goroutines and synchronize them using `sync.WaitGroup`. However, goroutines often need to **communicate** with each other. This is where **channels** come in—they provide a way to send and receive data between goroutines safely.  

#### **Key Concepts**  

1. **What Are Channels?**  
   - Channels allow goroutines to communicate safely without using shared memory.  
   - They work like **pipelines** where one goroutine sends data, and another receives it.  
   - Defined using `chan` keyword:  
     ```go
     ch := make(chan int)  // Creates a channel that carries int values
     ```

2. **Sending and Receiving Data**  
   - **Send:** `ch <- value` (sends a value into the channel)  
   - **Receive:** `val := <- ch` (receives a value from the channel)  
   - Channels block execution until both sender and receiver are ready.  

3. **Buffered vs. Unbuffered Channels**  
   - **Unbuffered channels** (default) block until the receiver is ready.  
   - **Buffered channels** allow storing multiple values before blocking.  
     ```go
     ch := make(chan int, 3) // Buffered channel with capacity 3
     ```

4. **Closing a Channel**  
   - `close(ch)` signals that no more values will be sent on the channel.  
   - Receivers can check if a channel is closed using:  
     ```go
     val, ok := <- ch // 'ok' is false if the channel is closed
     ```

---

### **Problem Statement**  
Create a program where the main function spawns multiple goroutines that send numbers into a channel. The main function should receive and print those numbers.  

---

### **Steps for Implementation**  

1. **Create an unbuffered channel to send integers.**  
2. **Launch multiple goroutines that send numbers into the channel.**  
3. **In the main function, receive values from the channel and print them.**  
4. **Close the channel once all data is sent to avoid deadlocks.**  
5. **Observe how channels block execution until data is received.**  