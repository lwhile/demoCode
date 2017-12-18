import json

class Animal(object):
    def __init__(self,name, traits):
        self.name = name 
        self.traits = traits

    def __str__(self):
        return self.name
 
def init_dataset():
    with open("data.json", "r") as f:
        return json.load(f)

def analysis(dataset):
    # 返回动物集合,以及所有动物的特征集合
    animals = []
    traits = set()
    for data in dataset:
        animal = Animal(data["result"], set(data["conditions"]))
        animals.append(animal)
        traits.update(set(data["conditions"]))
    return animals, traits

def query_user(traits):
    traits_from_user = set()
    for trait in traits:
        print("该动物是否有该特征:", trait)
        ans = input("输入y/n:")
        if ans == "y":
            traits_from_user.add(trait)
    return traits_from_user

def speculate_animal(animals,traits_from_user):
    ans_animals = []
    for animal in animals:
        inter_set = animal.traits & traits_from_user
        if len(inter_set) == len(animal.traits):
            ans_animals.append(animal)
    return ans_animals

if __name__ == "__main__" :
    dataset = init_dataset()
    animals, traits = analysis(dataset)
    traits_from_user = query_user(traits)
    ans_animals =  speculate_animal(animals, traits_from_user)
    if len(ans_animals) == 0:
        print("抱歉,找不到该动物.")
    else:
        for ans_animal in ans_animals:
            print(ans_animal.name)
    
        