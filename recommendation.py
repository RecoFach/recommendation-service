import pandas as pd
from numpy import dot
from numpy.linalg import norm
from typing import List, Tuple


def compare_to_subject_dataset(subject_user: pd.DataFrame,
                               subjects_dataset: pd.DataFrame) -> List[Tuple[int, int]]:
    """
    Compare the user settings to the dataset characteristics
    :param subject_user: vector of the user chosen fields
    :param subjects_dataset: dataset of the annotated subjects
    :return: a list of tuples for a subject index in the dataset and the score
    """
    weights_array = []
    for sample_train in range(subjects_dataset.shape[0]):
        b = subjects_dataset.iloc[sample_train, :]
        cos_sim = dot(subject_user, b) / (norm(subject_user) * norm(b))
        weights_array.append((sample_train, cos_sim))

    return weights_array


def get_title_from_index(df: pd.DataFrame,
                         index: int) -> pd.DataFrame:
    return df[df.index == index]["Course name"].values[0]


def get_index_from_title(df: pd.DataFrame,
                         title: str) -> pd.DataFrame:
    return df[df["Course name"] == title]["index"].values[0]


def recommend_subjects(subject_user: pd.DataFrame,
                       subjects_dataset: pd.DataFrame) -> List[Tuple[int, int]]:
    weights_array = compare_to_subject_dataset(subject_user=subject_user, subjects_dataset=subjects_dataset)
    sorted_similar_subjects = sorted(weights_array, key=lambda x: x[1], reverse=True)
    return sorted_similar_subjects


def main():
    df = pd.read_csv('subjects.csv', sep=";")

    user_query = {'Software engineering': [0], 'AI': [1], 'Low-level': [0], 'Security': [0], 'Web': [0],
                  'theoretical': [1]}
    # create a frame based on a dict values
    df_new = pd.DataFrame.from_dict(user_query)

    subframe_only_charachteristics = df[['Software engineering', 'AI', 'Low-level', 'Security', 'Web', 'theoretical']]

    sorted_similar_subjects = recommend_subjects(subject_user=df_new,
                                                 subjects_dataset=subframe_only_charachteristics)

    list_recommended_indexes = [recommended_index for recommended_index, _ in sorted_similar_subjects[:10]]
    for element in sorted_similar_subjects[:10]:
        print(get_title_from_index(df, element[0]))
    print(df.iloc[list_recommended_indexes])


if __name__ == '__main__':
    main()
