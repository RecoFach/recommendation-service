import pandas as pd
from numpy import dot
from numpy.linalg import norm
from typing import List, Tuple, Dict
from flask import request, Flask, Response, jsonify, abort
from pprint import pprint
import json

app = Flask(__name__)

def compare_to_subject_dataset(subject_user: pd.DataFrame,
                               subjects_dataset: pd.DataFrame) -> List[Tuple[int, int]]:
    """
    Compare the user settings to the dataset characteristics
    Here is a cosin similarity used to estimate how similar two vectors are

    Args:
        subject_user: vector of the user chosen fields
        subjects_dataset: dataset of the annotated subjects

    Returns:
        a list of tuples for a subject index in the dataset and the score
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
    """
    Call the cosin_similarity func and sort the results in descending order
    Args:
        subject_user: user query vector
        subjects_dataset: the vectors from subjects annoated in the csv

    Returns:
        Sorted based on the similarity subjects to the user_query

    """
    weights_array = compare_to_subject_dataset(subject_user=subject_user, subjects_dataset=subjects_dataset)
    sorted_similar_subjects = sorted(weights_array, key=lambda x: x[1], reverse=True)
    return sorted_similar_subjects


def process_user(user_query: Dict[str, List[int]], df: pd.DataFrame, amount_to_recommend: int = 5) -> None:
    """
    This function creates a recommendation based on the user_query.
    1. Select only the fields basde on which will be created the consin similarity
    2. Call the recommend_subjects func to get the most similar to the user_query data
    3. Use the top recommendated subjects and dump them to json

    Args:
        user_query: This is a vector based on which will be created recommendation of a subject f.e {'Software
    engineering': [0], 'AI': [1], 'Low-level': [0], 'Security': [0], 'Web': [0], 'theoretical': [1]} df: this is a 
    annotated file with subjects offered by the Technische Universität Dresden 

        amount_to_recommend: the amount of subjects that will be recommended to a user_query
    Returns:

    """
    df_new = pd.DataFrame.from_dict(user_query)
    subframe_only_charachteristics = df[['Software engineering', 'AI', 'Low-level', 'Security', 'Web', 'Theoretical']]
    sorted_similar_subjects = recommend_subjects(subject_user=df_new,
                                                 subjects_dataset=subframe_only_charachteristics)
    list_recommended_indexes = [recommended_index for recommended_index, _ in
                                sorted_similar_subjects[:amount_to_recommend]]

    recommended_subjects_frame = df.iloc[list_recommended_indexes].reset_index(drop=True)
    recommend_json = recommended_subjects_frame.to_json(orient="index")

    recommend_json_with_slashes = json.dumps(json.loads(recommend_json))
    return recommend_json_with_slashes

@app.route('/recommend', methods=['POST'])
def main():
    # subjects.csv contains the annotated data based on which will be created a similarity index
    df = pd.read_csv('subjects.csv', sep=";")
    data = request.get_json()
    user_query = {'Software engineering': [int(data.get('SOFTWARE'))],
                  'AI': [int(data.get('AI'))],
                  'Low-level': [int(data.get('LOWLEVEL'))],
                  'Security': [int(data.get('SECURITY'))],
                  'Web': [int(data.get('WEB'))],
                  'Theoretical': [int(data.get('THEORETICAL'))]}

    # processing the user_data and create an recommendation
    preddicted_recoomendation = process_user(user_query, df)
    return preddicted_recoomendation


if __name__ == '__main__':
    app.run('0.0.0.0', port=2000)
