# Debounce
A small Go library for debouncing noisy channels.

### Example
```golang
noisy := make(chan interface{})
go func() {
  for {
    noisy <- interface{}(rand.Intn(2)) // Random signals every millisecond
    time.Sleep(time.Millisecond)
  }
}()

debounced := debounce.Channel(noisy, 10*time.Millisecond)
for stable := range debounced {
  // Signal only if the value has not changed for 10 milliseconds
}
```
