package common

type ConfigParser func(sectionReader SectionReader, text string) SectionalConfiguration
