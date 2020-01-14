/*
 * person.h
 */
#ifndef __PERSON_H__
#define __PERSON_H__

typedef struct APerson {
    const char *name;
    const char *long_name;
} APerson;

APerson *get_person(const char *name, const char *long_name);

#endif // __PERSON_H__
