#include "iostream"
#include "vector"

void Swap(int* a, int* b);
void BubbleSort(std::vector<int> &array);
std::vector<int> BubbleSorted(std::vector<int> arr);
void PrintArray(std::vector<int> array);

void BubbleSort (std::vector<int> &array)
{
    std::cout<<"Elements in the array: "<<array.size()<<std::endl;

    for (int i = 0; i < array.size(); i++)
    {
        for(int j = 0; j < array.size() - 1; j++)
        {   
            if (array[j] > array[j+1])
                Swap(&array[j], &array[j+1]);
        }
    }
}


void PrintArray (std::vector<int> array)
{
    for (int i = 0; i < array.size(); i++)
        std::cout<<array[i]<<" ";
    std::cout<<std::endl;
}


std::vector<int> BubbleSorted(std::vector<int> arr){
	std::vector<int> temp = arr;
	BubbleSort(&temp);

	return temp;
}

void Swap (int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}
