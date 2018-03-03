# Alien Invasion

Input is 
* city_map a file that is a map of cities, and
* N the number of aliens to prepopulate in the cities

## Main Flows of Code
Here is my First Draft of Logic
```
isDone = false

while !isDone:

    // Will Only Switch to False if At Least One Alien Alive
    // That Hasn't Moved 10,000 times
    isDone = true

    for a in aliens:
        if a.isAlive():
            a.TravelToCity()
            if a.NumMoves < 10000:
                isDone = false

    for c in cities:   
        if c.NumAliens() > 1:
            c.Explode()

// Print Cities
for c in cities:
    print c.String()

```
Further inspection reveals `a.TravelToCity()` might be something like,
```
self.leave(prevCity) // Deregister from old city
self.arrive(newCity) // Register to new city
```
Further inspection of `c.Explode()` might look like,
```
// Pick Which Two Aliens Fought
leftAlien = self.aliens[0]
rightAlien = self.aliens[1]

// Print Destruction Message
self.AnnounceDestroyed()

// Destroy Roads In and Out
for n in self.neighboringCities:
    n.DestroyRoadTo(self)
    n = nil

// Kill all Aliens in City
for a in self.aliens:
    a.Kill()
```