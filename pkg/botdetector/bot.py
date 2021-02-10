import pandas

from sklearn.tree import DecisionTreeClassifier
from sklearn.metrics import accuracy_score
from sklearn.model_selection import train_test_split
from os import path


def main():
    # parse the training data
    data = pandas.read_csv('/Users/tsmith/Documents/Projects/Social-Data-Collector/pkg/botdetector/Cleaned_TrainData.csv', encoding = "ISO-8859-1")

    # clean data
    data['screen_name_binary'] = data.screen_name.str.contains("", case=False, na=False)
    data['name_binary'] = data.name.str.contains("", case=False, na=False)
    data['description_binary'] = data.description.str.contains("", case=False, na=False)
    data['status_binary'] = data.status.str.contains("", case=False, na=False)
    data['listedcount_binary'] = (data.listedcount>20000)==False

    # declare features
    features = ['screen_name_binary', 'name_binary', 'description_binary', 'status_binary', 'verified', 'followers_count', 'friends_count', 'statuses_count', 'listedcount_binary', 'bot']

    # feature selection stage
    X = data[features].iloc[:,:-1] # independant feature variables (twitter account attributes)
    y = data[features].iloc[:,-1] # dependant target variable (bot indication)

    # split dataset into training and test sets
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.5, random_state=110)

    # preform predictions using Decision Tree classifier
    DecisionTree(X_train, X_test, y_train, y_test)

def DecisionTree(X_train, X_test, y_train, y_test):
    # create decision tree classifier
    # train the classifier given the training portion of the dataset
    clf = DecisionTreeClassifier(criterion='entropy', min_samples_leaf=50, min_samples_split=10)
    clf.fit(X_train, y_train)

    # use the classifier to predict out dependant bot target variable on the training, and then test dataset
    y_pred_train = clf.predict(X_train)
    y_pred_test = clf.predict(X_test)

    # use the trained model and the independant features training data to predict if a user is a bot
    prediction = predictUser(clf)

    # get training accuracy
    training_accuracy = str(round(accuracy_score(y_train, y_pred_train) * 100)) + "%"


    if prediction[0] == 1:
        print("BOT-" +training_accuracy)
        return
    print("HUMAN-" +training_accuracy)
    return

# predictUser uses the trained model and the features from the provided user.csv file to predict if that user is a bot
def predictUser(clf):
    independant_features = getCleanedFeaturesFromFile(path.abspath("../pkg/botdetector/user.csv"))
    prediction = clf.predict(independant_features)
    return prediction



def getCleanedFeaturesFromFile(path):
    # parse the training data
    data = pandas.read_csv(path, encoding = "ISO-8859-1")

    # clean data
    data['screen_name_binary'] = data.screen_name.str.contains("", case=False, na=False)
    data['name_binary'] = data.name.str.contains("", case=False, na=False)
    data['description_binary'] = data.description.str.contains("", case=False, na=False)
    data['status_binary'] = data.status.str.contains("", case=False, na=False)
    data['listedcount_binary'] = (data.listedcount>20000)==False

    # declare features
    features = ['screen_name_binary', 'name_binary', 'description_binary', 'status_binary', 'verified', 'followers_count', 'friends_count', 'statuses_count', 'listedcount_binary', 'bot']

    # feature selection stage
    independantFeatures = data[features].iloc[:,:-1] # independant feature variables (twitter account attributes)
    return independantFeatures


if __name__ == '__main__':
    main()