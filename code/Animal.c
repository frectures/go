#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Animal {
    // polymorphism
    struct Animal_vtable {
        const char* (*noise)(const struct Animal*);
    };
    const struct Animal_vtable* vptr; // configured by "constructors"

    char name[16];
};

//////////////////////////////////////////// abstract
const struct Animal_vtable animal_vtable = { NULL };

// constructor
void Animal_init(struct Animal* animal, const char* name) {
    // polymorphism
    animal->vptr = &animal_vtable;

    strncpy(animal->name, name, 15);
    animal->name[15] = 0;
};

void chat(struct Animal* animal) {
    ///////////////////////////////////////////// polymorphism
    printf("%s says: %s\n", animal->name, animal->vptr->noise(animal));
}

typedef enum { false, true } bool;

struct Cow {
    // inheritance
    struct Animal base;

    bool annoyed;
};

const char* Cow_noise(const struct Animal* animal) {
    const struct Cow* cow = (const struct Cow*) animal;
    return cow->annoyed ? "MOOOOO!!!" : "Mooooo...";
}

///////////////////////////////////////// override
const struct Animal_vtable cow_vtable = { &Cow_noise };

// constructor
void Cow_init(struct Cow* cow, const char* name) {
    // chaining
    Animal_init(&cow->base, name);
    // polymorphism
    cow->base.vptr = &cow_vtable;

    cow->annoyed = (rand() % 3) < 1;
}

int main() {
    struct Cow cow;
    Cow_init(&cow, "Emma");
    // polymorphism
    struct Animal* animal = (struct Animal*) &cow;
    chat(animal);
}
