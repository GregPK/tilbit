# TILBit

Retain the information you consume. Think deeply about things you learn.

## Assumptions

- Quality of information is more important then quantity.
- You can only retain quality information by reflecting on it periodically.
- Timing of information is very important, but hard to pin down. The best we can do is to trigger recall often and hope it's relevant.

## Features

- Quickly add thoughts throught the command line. Supports in-line JSON metadata.
- Parse highlights done on a Kindle (via "My Clippings.txt" file)
- Support for a custom Markdown file format for TILs from one source.

## Usage

Get random tilbit from [default sources](#default_sources):

```bash
tilbit
# or
tilbit show random
```

Get a specific tilbit:

```bash
tilbit show 1a4f5080
```

Get all tilbits:

```bash
tilbit show all
```

Load from a specific source:

```bash
tilbit show random --source=seeds/  # for directory
tilbit show random --source=seeds/ralph-waldo-emerson.tilbit.md  # for specific file
```

Show specific [format](#formats):

```bash
tilbit show random -f=box
tilbit show random -f=text
tilbit show random -f=yaml
```

### Formats

- `box` - default terminal UI like decorator
- `text` - plaintext, with source data in specific format
- `yaml` - plaintext YAML, dumping all the data about an item, useful for inspection

### Configuration

#### Default sources

## Suggestions for use

- Do not fill your database with quotes from random sources. Only add things that really spoke to you.
- Focus on your own learnings, rather than "inspiration". The goal is to make you think about the media you consume and consciously learn from it.
- Do not drown yourself with recall. Aim for one item per hour *at most*. Think about the item for *at least* 30 seconds.
- Focus in "write" rather than "read", especially in the beginning. You're instinct will be to be a consumer of data, we want to transform ourselves into being able to think crititally about what we are consuming and thinking about it deeply.
- A negative result is a result. If you haven't learned anything worthwhile, think about how you got to that particular item. What made you want to read it? How did it draw you in.


## Temporary database migration plan - to be removed.

I want to transition the tool to have a permanent local database of items. This is general has a lot of strong points, mostly relating to how data can be queried without re-inventing the wheel.

To keep the simplicity of having just a set of text files that can be freely edited, I've decided to keep the option of having them as sources. The db will be used to annotate them. The only problem this approach has, is that if the content of the tilbit changes, there will be orphaned items in the database. This could potentially be solved by caching the content of the tilbit and if any orpans are detected, to try to find the new iteration using some kind of fuzzy text metric (like Levenshtein distance).

I want to do this in a very non-breaking and small steps approach so the tentative plan is to:

- [x] keep the source approach for now
- [x] refactor the current sources into a new repository struct
- [x] implement and test the db handlers and all the db as a new alternative repository
- [x] implement importers that will wrap current parsers
- [x] the db will be an in-memory database for the time being that runs importers on each run
- [x] this will allow for gradual testing and moving to a new approach
- [ ] at some point, I will just move the sqlite from in-memory to a file and cut down on the overhead
