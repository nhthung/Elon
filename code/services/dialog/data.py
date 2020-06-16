import os
import pandas as pd

columns = [
    'intent',
    'train_phrases',
    'responses',
    'entities',
    'contexts',
    'followup_intents'
]


def process(data_dir, filename, sheet_name, columns, skiprows=0):
    path = f'{data_dir}/{filename}'

    mapping = (
        ('train_phrases', sanitize_train_phrases),
        ('responses', sanitize_responses)
    )

    if not os.path.isfile(path):
        raise ValueError('File not found')
    
    data = pd.read_excel(path, sheet_name, skiprows)
    # data = data.iloc[:, :len(columns)]
    data.columns = columns

    for col, func in mapping:
        data.loc[:, col] = data.loc[:, col].map(func)

    return data


def sanitize_train_phrases(train_phrases):
    if isinstance(train_phrases, str):
        return tuple(map(sanitize_train_phrase, train_phrases.strip().split('\n')))
    else:
        return ()


def sanitize_train_phrase(train_phrase):
    return re.sub(r'[\?!\.]', '', train_phrase.lower())


def sanitize_responses(responses):
    if isinstance(responses, str):
        return tuple(response.strip().split('\n') for response in responses.split('&&'))
    else:
        return ()