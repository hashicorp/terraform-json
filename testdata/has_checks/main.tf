module "files" {
  source     = "./child_module"
  file_names = ["file1.txt", "file2.txt"]
}