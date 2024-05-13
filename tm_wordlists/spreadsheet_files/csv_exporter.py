# ignore_for_file: avoid_print

import os
import pandas as pd


class CsvEntry:
    def __init__(self, category, subcategory, localTerm, englishTranslation):
        self.category = category
        self.subcategory = subcategory
        self.localTerm = localTerm
        self.englishTranslation = englishTranslation


def main():
    directory = os.getcwd()

    # iterate through .xslx files in ./bin/
    for filename in os.listdir(directory):
        if not filename.endswith('.xlsx'):
            continue

        filepath = os.path.join(directory, filename)
        excel = pd.read_excel(filepath, sheet_name=None)

        entries = []
        try:
            if len(excel.keys()) == 1:
                for sheetKey, sheet in excel.items():
                    for rowIndex in range(0, len(sheet)):
                        entries.append(
                            CsvEntry(sheet.iloc[rowIndex, 0], sheet.iloc[rowIndex, 1], sheet.iloc[rowIndex, 2],
                                     sheet.iloc[rowIndex, 3]))
            else:
                for sheetKey, sheet in excel.items():
                    category = sheetKey
                    # find indices of Columns that have first row containing 'Subcategory'
                    for columnIndex in range(len(sheet.columns)):
                        if sheet.columns[columnIndex] == 'Subcategory':
                            subcategoryColumnIndex = columnIndex
                            subcategory = sheet.iloc[0, subcategoryColumnIndex]

                            if pd.isnull(subcategory):
                                raise Exception('Subcategory is empty')

                            localTermColumnIndex = columnIndex + 2
                            englishTranslationColumnIndex = localTermColumnIndex + 1

                            for rowIndex in range(0, len(sheet)):
                                localTerm = sheet.iloc[rowIndex, localTermColumnIndex]
                                englishTranslation = sheet.iloc[rowIndex, englishTranslationColumnIndex]

                                if not pd.isnull(localTerm) and not pd.isnull(englishTranslation):
                                    entries.append(
                                        CsvEntry(
                                            category,
                                            subcategory,
                                            str(localTerm),
                                            str(englishTranslation),
                                        )
                                    )

            # write entries to a .csv file in the format:
            # Category,Subcategory,Local term,English translation
            csvFileName = filename.split('.')[0] + ".csv"
            csvFilePath = os.path.join(directory, csvFileName)

            with open(csvFilePath, 'w') as csvFile:
                csvFile.write('Category,Subcategory,Local term,English translation,Simple category\n')

                for entry in entries:
                    csvFile.write(
                        f'{entry.category},{entry.subcategory},{entry.localTerm},{entry.englishTranslation},{entry.subcategory}\n'
                    )

                csvFile.close()

            print(f'Finished - {filepath} 🎉 Wrote {len(entries)} rows of data')
        except Exception as e:
            print('Error')
            print(e)


if __name__ == "__main__":
    main()
