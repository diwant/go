# Alien Invasion

To run, navigate to the alien_invasion directory, and execute,

```
go build 
./alien_invasion cities.txt 2
```
Here, `cities.txt` can be replaced by the file describing your cities, and `2` can be replaced by the number of aliens in your world.

Once again, that input is 
* city_map a file that is a map of cities, and
* N the number of aliens to prepopulate in the cities

## Assumptions Made
+ Links between cities are as provided.  So just because Foo has a road to Baz, Baz does not have a road to Foo unless explicitly declared in the input file.  

+ Multiple aliens can be in any given city at the end of an alien migration.  The first two on the city's registry will fight, the city explodes, and all aliens in the city die, not just the ones fighting.

+ Moving one city is one turn for an alien.  Moving Through a city to another beyond requires more than 1 turn and the alien might encounter in the middle another alien who it must then battle

## Logic Draft
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